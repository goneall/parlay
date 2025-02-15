/*
 * © 2023 Snyk Limited All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ecosystems

import (
	"errors"
	"strings"

	"github.com/package-url/packageurl-go"
	"github.com/rs/zerolog"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/spdx/v2/common"
	"github.com/spdx/tools-golang/spdx/v2/v2_3"

	"github.com/snyk/parlay/ecosystems/packages"
	"github.com/snyk/parlay/internal/utils"
)

func enrichSPDX(bom *spdx.Document, logger *zerolog.Logger) {
	bomPackages := bom.Packages

	logger.Debug().Msgf("Detected %d packages", len(bomPackages))

	pkgDataCache := make(map[string]*packages.Package)
	pgkVersionDataCache := make(map[string]*packages.VersionWithDependencies)

	for _, pkg := range bomPackages {
		purl, err := extractPurl(pkg)
		if err != nil {
			continue
		}

		pkgData, ok := pkgDataCache[purl.String()]

		if !ok {
			packageResp, err := GetPackageData(*purl)
			if err != nil {
				continue
			}

			pkgData = packageResp.JSON200
			pkgDataCache[purl.String()] = pkgData
		}

		if pkgData == nil {
			continue
		}
		enrichSPDXDescription(pkg, pkgData)
		enrichSPDXHomepage(pkg, pkgData)
		enrichSPDXSupplier(pkg, pkgData)

		pkgVersionData, ok := pgkVersionDataCache[purl.String()]
		if !ok {
			packageVersionResp, err := GetPackageVersionData(*purl)
			if err != nil {
				continue
			}

			pkgVersionData := packageVersionResp.JSON200
			pgkVersionDataCache[purl.String()] = pkgVersionData
		}
		if pkgVersionData == nil {
			continue
		}

		enrichSPDXLicense(pkg, pkgVersionData, pkgData)
	}
}

func extractPurl(pkg *v2_3.Package) (*packageurl.PackageURL, error) {
	for _, ref := range pkg.PackageExternalReferences {
		if ref.RefType != "purl" {
			continue
		}
		purl, err := packageurl.FromString(ref.Locator)
		if err != nil {
			return nil, err
		}
		return &purl, nil
	}
	return nil, errors.New("no purl found on SPDX package")
}

func enrichSPDXSupplier(pkg *v2_3.Package, data *packages.Package) {
	if data.RepoMetadata != nil {
		meta := *data.RepoMetadata
		if ownerRecord, ok := meta["owner_record"].(map[string]interface{}); ok {
			if name, ok := ownerRecord["name"].(string); ok && name != "" {
				pkg.PackageSupplier = &common.Supplier{
					SupplierType: "Organization",
					Supplier:     name,
				}
			}
		}
	}
}

func enrichSPDXLicense(pkg *v2_3.Package, pkgVersionData *packages.VersionWithDependencies, pkgData *packages.Package) {
	licenses := utils.GetLicensesFromEcosystemsLicense(pkgVersionData, pkgData)
	if len(licenses) > 0 {
		pkg.PackageLicenseConcluded = strings.Join(licenses, ",")
	}
}

func enrichSPDXHomepage(pkg *v2_3.Package, data *packages.Package) {
	if data.Homepage == nil {
		return
	}
	pkg.PackageHomePage = *data.Homepage
}

func enrichSPDXDescription(pkg *v2_3.Package, data *packages.Package) {
	if data.Description == nil {
		return
	}
	pkg.PackageDescription = *data.Description
}
