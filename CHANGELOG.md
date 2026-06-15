# Changelog

## [1.12.0](https://github.com/michaeldcanady/servicenow-sdk-go/compare/v1.11.1...v1.12.0) (2026-06-15)


### ⚠ BREAKING CHANGES

* Removed incorrect methods `CreateOrUpdateService` and `GetContent` from `AppServiceRequestBuilder` and nested items.
* Renames CmdbInstanceRequestBuilder2 to CmdbInstanceRequestBuilder. Changes Post, Put, and Patch method signatures to accept the request body as a separate parameter.
* Renames CmdbInstanceRequestBuilder2 to CmdbInstanceRequestBuilder. Changes Post, Put, and Patch method signatures to accept the request body as a separate parameter.

### Features

* **actsub:** implement ServiceNow Activity Subscriptions API support ([d377fb1](https://github.com/michaeldcanady/servicenow-sdk-go/commit/d377fb10acef21947f124f728c477c649bfcdbd9))
* add account api ([#444](https://github.com/michaeldcanady/servicenow-sdk-go/issues/444)) ([0c9f5b3](https://github.com/michaeldcanady/servicenow-sdk-go/commit/0c9f5b37dfdfae9c4e073fb834d0b723c414c112))
* add account api ([#446](https://github.com/michaeldcanady/servicenow-sdk-go/issues/446)) ([9833e5d](https://github.com/michaeldcanady/servicenow-sdk-go/commit/9833e5dea15b8bb5bc8eb183d1af7d70a1329052))
* add appointment api ([#447](https://github.com/michaeldcanady/servicenow-sdk-go/issues/447)) ([b912091](https://github.com/michaeldcanady/servicenow-sdk-go/commit/b912091cf790d5480b854e36e10f3f032644f877))
* add case api support ([#448](https://github.com/michaeldcanady/servicenow-sdk-go/issues/448)) ([174fc91](https://github.com/michaeldcanady/servicenow-sdk-go/commit/174fc91267540943c70021efab7e46743f3289fb))
* add cdm editor api module ([#449](https://github.com/michaeldcanady/servicenow-sdk-go/issues/449)) ([d65bb0c](https://github.com/michaeldcanady/servicenow-sdk-go/commit/d65bb0cc4369a98e14eb745e233ec7e9ed22e725))
* add ServiceNow CMDB Instance API support ([#443](https://github.com/michaeldcanady/servicenow-sdk-go/issues/443)) ([b9dc6be](https://github.com/michaeldcanady/servicenow-sdk-go/commit/b9dc6be942cf91b81779df114b2125660afbe5b4))
* add ServiceNow Documents API support ([#442](https://github.com/michaeldcanady/servicenow-sdk-go/issues/442)) ([4a97413](https://github.com/michaeldcanady/servicenow-sdk-go/commit/4a9741376d3971ac7969cdefe0d60a6bb1d04547))
* add specification for prerelease workflow ([b4fbc0e](https://github.com/michaeldcanady/servicenow-sdk-go/commit/b4fbc0ed389b1ebd81dc905083035cd70dc29aa6))
* add support for application service api ([#452](https://github.com/michaeldcanady/servicenow-sdk-go/issues/452)) ([f236120](https://github.com/michaeldcanady/servicenow-sdk-go/commit/f23612081762ff8b6d8b8086ff14470e6a5d710e))
* add support for cdm applications api ([#451](https://github.com/michaeldcanady/servicenow-sdk-go/issues/451)) ([8b4ab42](https://github.com/michaeldcanady/servicenow-sdk-go/commit/8b4ab42f009069a975a96ecae49ada8e799ad49a))
* add support for cdm changeset api ([#450](https://github.com/michaeldcanady/servicenow-sdk-go/issues/450)) ([6bae5c5](https://github.com/michaeldcanady/servicenow-sdk-go/commit/6bae5c55d278a697c38dd4ad34d95d7c929db62c))
* integrate AI-driven API generation framework ([#440](https://github.com/michaeldcanady/servicenow-sdk-go/issues/440)) ([59927ed](https://github.com/michaeldcanady/servicenow-sdk-go/commit/59927ed2073f86d2baacdb86723ed05662dfbf1f))


### Bug Fixes

* **api:** correct model deserialization and remove trailing newline ([efab807](https://github.com/michaeldcanady/servicenow-sdk-go/commit/efab80796d402abb4655bd76db64f025d7c6ef4f))
* **document:** remove unneeded "2" suffix ([b64dcec](https://github.com/michaeldcanady/servicenow-sdk-go/commit/b64dcec24811ace4ad8e8fdfe6ee35be52584378))
* fix broken NowRequestBuilder2 ([9f3db3a](https://github.com/michaeldcanady/servicenow-sdk-go/commit/9f3db3a3f66f9774e98ceb2f8547a1a13d44cc99))
* remove problematic commented-out features from devcontainer.json ([e8d7b27](https://github.com/michaeldcanady/servicenow-sdk-go/commit/e8d7b274d2fdad160504e3eea03b88d70fb61984))
* resolve dev container boot failure by hardcoding mount paths ([c004deb](https://github.com/michaeldcanady/servicenow-sdk-go/commit/c004deb65c2e86093f41f30baced2543020c5c57))

## [1.11.1](https://github.com/michaeldcanady/servicenow-sdk-go/compare/v1.11.0...v1.11.1) (2026-03-28)


### Bug Fixes

* **credential:** fix "unsupported protocol scheme" ([#427](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/427)) ([a616770](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/a61677089c5c494d0003f0a99ddfc29f3dd98616))

## [1.11.0](https://github.com/michaeldcanady/servicenow-sdk-go/v2/compare/v1.10.0...v1.11.0) (2026-03-22)


### Features

* add bug-reporter gemini skill ([#340](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/340)) ([f91fa1b](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/f91fa1b935040502a0bb29cc0c4d33dcd7a06bfe))
* add support for ropc credentials ([01d7365](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/01d7365e142c1ec48606876de98f9890d79bab73))
* **credential:** add support for local redirect server timeout ([3db464e](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/3db464ed6751a1abda6724c3511d4085457e4cef))
* **credential:** auth code provider ([#417](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/417)) ([5f5d6b2](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/5f5d6b243fe89c674eff6bcab3ea09e0df6e7c90))
* **credential:** client credentials provider ([#416](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/416)) ([18c9de8](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/18c9de8f2b6db49b66fbe234a37ddedee0fcbd0d))
* **credential:** make authorization code credential thread-safe ([4d710e1](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/4d710e1cd806b108c3002813c4e6e9abe679e58a))
* **credentials:** jwt grant ([#421](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/421)) ([fb9b5a5](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/fb9b5a5b468c48edc7986a40b2448d775db6756c))
* **internal:** add request configuration helper and update page iterator ([7fcbf53](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/7fcbf534367210e8850c6e5b52fa77d1c8922836))
* make data accessible ([7471d49](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/7471d4955a98fa838443accaa672dbb127faa380))
* **policy api:** now policies api implement list definitions ([#409](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/409)) ([1c563de](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/1c563deaa69ee5c5f212ecfe968f8a9a2dd7cf18))
* **policy:** delete mapping ([#410](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/410)) ([5940134](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/5940134b08e4a208f13482d4369b2f0d18078735))
* **policy:** scaffold list definitions request builder ([#386](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/386)) ([#406](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/406)) ([f946bf3](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/f946bf3a8cb8b441efc191eae3e21120325f42fc))
* Reimplementing ropc credentials ([#334](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/334)) ([ecc16ae](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/ecc16aed98a9c622d48fdc007d135bd4c61382ab))


### Bug Fixes

* (credential): fix client secret not being added to auth url ([1aee930](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/1aee93073484cb3fcf5b9ae7f642bdc46ff3138a))
* add missing release-please version tags ([#328](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/328)) ([fa83ff1](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/fa83ff17cc09bb7a93091f8338d3f1693cef06f5))
* add serialization to table record ([aeeb45e](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/aeeb45e4d0f9cd591ec9d3909785acb6b8b09371))
* allow serialization ([29be0fd](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/29be0fd1503f2936d3283d03963fd879bd6535bd))
* bad ref name ([#331](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/331)) ([d64cb4a](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/d64cb4ae95f6d8b0ec966e7798f88bc670027009))
* changelog preventing go get ([#420](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/420)) ([c6c0ce6](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/c6c0ce6716e59216c3aee97d6dc208a5ea980b46))
* checkout code prior to updating unreleased ([#330](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/330)) ([714b67a](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/714b67afd0f3893718fabc63396e7cfe0efd8c70))
* **ci:** use setup-go built-in caching to prevent cache pollution ([#341](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/341)) ([f837777](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/f837777f003fda850b102ba312cdba1c6e6e27a4)), closes [#339](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/339)
* correct key capitalization ([974c125](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/974c125dab83996fa2f96aff4d520ae0de1c9104))
* merge error ([13d1eb5](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/13d1eb5c41c7afc2f4c10f16861007f7a57affd9))
* not include component name in tag ([83caa5e](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/83caa5e08fb06574e44793871a8c539b87357334))
* **paging:** ensure Link headers are correctly captured and parsed ([58e43b6](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/58e43b6706eef8b27df99b9891d447fe042f6ef3))
* **REFRESH_TOKEN:** Fixed a trailing whitespace in the refresh_token parameter ([#303](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/303)) ([a88a4b2](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/a88a4b27013a7788b4fa0d0023ade64552b67bee))
* remove "v" prefix ([8a5b940](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/8a5b940bad3d61522ce041dba8ecbcc8c6c112af))
* resolve integration test failures and SDK URL bugs ([#345](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/345)) ([d8c4afc](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/d8c4afc62eacceaa7dd1bed70d1f55e36586a015))
* strip unneeded comment ([71af4f9](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/71af4f9bcb98d088c3daf6f88a13a5f5b11faf3d))
* unable to get response from new table api implementation ([ce520ba](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/ce520ba12946040e557f669abef773e454590842))

## [1.10.0](https://github.com/michaeldcanady/servicenow-sdk-go/v2/compare/github.com/michaeldcanady/servicenow-sdk-go/v2-v1.9.0...github.com/michaeldcanady/servicenow-sdk-go/v2-v1.10.0) (2026-03-06)


### Features

* add bug-reporter gemini skill ([#340](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/340)) ([f91fa1b](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/f91fa1b935040502a0bb29cc0c4d33dcd7a06bfe))
* add support for ropc credentials ([01d7365](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/01d7365e142c1ec48606876de98f9890d79bab73))
* **internal:** add request configuration helper and update page iterator ([7fcbf53](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/7fcbf534367210e8850c6e5b52fa77d1c8922836))
* make data accessible ([7471d49](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/7471d4955a98fa838443accaa672dbb127faa380))
* Reimplementing ropc credentials ([#334](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/334)) ([ecc16ae](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/ecc16aed98a9c622d48fdc007d135bd4c61382ab))


### Bug Fixes

* add serialization to table record ([aeeb45e](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/aeeb45e4d0f9cd591ec9d3909785acb6b8b09371))
* allow serialization ([29be0fd](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/29be0fd1503f2936d3283d03963fd879bd6535bd))
* **ci:** use setup-go built-in caching to prevent cache pollution ([#341](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/341)) ([f837777](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/f837777f003fda850b102ba312cdba1c6e6e27a4)), closes [#339](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/339)
* correct key capitalization ([974c125](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/974c125dab83996fa2f96aff4d520ae0de1c9104))
* **paging:** ensure Link headers are correctly captured and parsed ([58e43b6](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/58e43b6706eef8b27df99b9891d447fe042f6ef3))
* resolve integration test failures and SDK URL bugs ([#345](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/345)) ([d8c4afc](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/d8c4afc62eacceaa7dd1bed70d1f55e36586a015))
* unable to get response from new table api implementation ([ce520ba](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/ce520ba12946040e557f669abef773e454590842))

## [1.9.0](https://github.com/michaeldcanady/servicenow-sdk-go/v2/compare/github.com/michaeldcanady/servicenow-sdk-go/v2-v1.8.0...github.com/michaeldcanady/servicenow-sdk-go/v2-v1.9.0) (2026-03-02)


### Features

* added Last method ([c7b440f](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/c7b440f9ccaa2638a5151e168b37ff58258757f9))
* added Last method ([f01d7a0](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/f01d7a0c5dd06feac1c21aa3b0c03edde36acd84))
* added Last method ([846052c](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/846052ce27f7717811d514d993bcec9d55a28704))
* added Last method ([28f6ccc](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/28f6ccc35444711084b4131864b64c05ca9e16f3))


### Bug Fixes

* add missing release-please version tags ([#328](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/328)) ([fa83ff1](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/fa83ff17cc09bb7a93091f8338d3f1693cef06f5))
* bad ref name ([#331](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/331)) ([d64cb4a](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/d64cb4ae95f6d8b0ec966e7798f88bc670027009))
* checkout code prior to updating unreleased ([#330](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/330)) ([714b67a](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/714b67afd0f3893718fabc63396e7cfe0efd8c70))
* **REFRESH_TOKEN:** Fixed a trailing whitespace in the refresh_token parameter ([#303](https://github.com/michaeldcanady/servicenow-sdk-go/v2/issues/303)) ([a88a4b2](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/a88a4b27013a7788b4fa0d0023ade64552b67bee))
* remove "v" prefix ([8a5b940](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/8a5b940bad3d61522ce041dba8ecbcc8c6c112af))
* strip unneeded comment ([71af4f9](https://github.com/michaeldcanady/servicenow-sdk-go/v2/commit/71af4f9bcb98d088c3daf6f88a13a5f5b11faf3d))
