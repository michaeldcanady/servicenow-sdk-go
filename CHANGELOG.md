# Changelog

## [1.11.0](https://github.com/michaeldcanady/servicenow-sdk-go/compare/v1.10.0...v1.11.0) (2026-03-22)


### Features

* add bug-reporter gemini skill ([#340](https://github.com/michaeldcanady/servicenow-sdk-go/issues/340)) ([f91fa1b](https://github.com/michaeldcanady/servicenow-sdk-go/commit/f91fa1b935040502a0bb29cc0c4d33dcd7a06bfe))
* add support for ropc credentials ([01d7365](https://github.com/michaeldcanady/servicenow-sdk-go/commit/01d7365e142c1ec48606876de98f9890d79bab73))
* **credential:** add support for local redirect server timeout ([3db464e](https://github.com/michaeldcanady/servicenow-sdk-go/commit/3db464ed6751a1abda6724c3511d4085457e4cef))
* **credential:** auth code provider ([#417](https://github.com/michaeldcanady/servicenow-sdk-go/issues/417)) ([5f5d6b2](https://github.com/michaeldcanady/servicenow-sdk-go/commit/5f5d6b243fe89c674eff6bcab3ea09e0df6e7c90))
* **credential:** client credentials provider ([#416](https://github.com/michaeldcanady/servicenow-sdk-go/issues/416)) ([18c9de8](https://github.com/michaeldcanady/servicenow-sdk-go/commit/18c9de8f2b6db49b66fbe234a37ddedee0fcbd0d))
* **credential:** make authorization code credential thread-safe ([4d710e1](https://github.com/michaeldcanady/servicenow-sdk-go/commit/4d710e1cd806b108c3002813c4e6e9abe679e58a))
* **credentials:** jwt grant ([#421](https://github.com/michaeldcanady/servicenow-sdk-go/issues/421)) ([fb9b5a5](https://github.com/michaeldcanady/servicenow-sdk-go/commit/fb9b5a5b468c48edc7986a40b2448d775db6756c))
* **internal:** add request configuration helper and update page iterator ([7fcbf53](https://github.com/michaeldcanady/servicenow-sdk-go/commit/7fcbf534367210e8850c6e5b52fa77d1c8922836))
* make data accessible ([7471d49](https://github.com/michaeldcanady/servicenow-sdk-go/commit/7471d4955a98fa838443accaa672dbb127faa380))
* **policy api:** now policies api implement list definitions ([#409](https://github.com/michaeldcanady/servicenow-sdk-go/issues/409)) ([1c563de](https://github.com/michaeldcanady/servicenow-sdk-go/commit/1c563deaa69ee5c5f212ecfe968f8a9a2dd7cf18))
* **policy:** delete mapping ([#410](https://github.com/michaeldcanady/servicenow-sdk-go/issues/410)) ([5940134](https://github.com/michaeldcanady/servicenow-sdk-go/commit/5940134b08e4a208f13482d4369b2f0d18078735))
* **policy:** scaffold list definitions request builder ([#386](https://github.com/michaeldcanady/servicenow-sdk-go/issues/386)) ([#406](https://github.com/michaeldcanady/servicenow-sdk-go/issues/406)) ([f946bf3](https://github.com/michaeldcanady/servicenow-sdk-go/commit/f946bf3a8cb8b441efc191eae3e21120325f42fc))
* Reimplementing ropc credentials ([#334](https://github.com/michaeldcanady/servicenow-sdk-go/issues/334)) ([ecc16ae](https://github.com/michaeldcanady/servicenow-sdk-go/commit/ecc16aed98a9c622d48fdc007d135bd4c61382ab))


### Bug Fixes

* (credential): fix client secret not being added to auth url ([1aee930](https://github.com/michaeldcanady/servicenow-sdk-go/commit/1aee93073484cb3fcf5b9ae7f642bdc46ff3138a))
* add missing release-please version tags ([#328](https://github.com/michaeldcanady/servicenow-sdk-go/issues/328)) ([fa83ff1](https://github.com/michaeldcanady/servicenow-sdk-go/commit/fa83ff17cc09bb7a93091f8338d3f1693cef06f5))
* add serialization to table record ([aeeb45e](https://github.com/michaeldcanady/servicenow-sdk-go/commit/aeeb45e4d0f9cd591ec9d3909785acb6b8b09371))
* allow serialization ([29be0fd](https://github.com/michaeldcanady/servicenow-sdk-go/commit/29be0fd1503f2936d3283d03963fd879bd6535bd))
* bad ref name ([#331](https://github.com/michaeldcanady/servicenow-sdk-go/issues/331)) ([d64cb4a](https://github.com/michaeldcanady/servicenow-sdk-go/commit/d64cb4ae95f6d8b0ec966e7798f88bc670027009))
* changelog preventing go get ([#420](https://github.com/michaeldcanady/servicenow-sdk-go/issues/420)) ([c6c0ce6](https://github.com/michaeldcanady/servicenow-sdk-go/commit/c6c0ce6716e59216c3aee97d6dc208a5ea980b46))
* checkout code prior to updating unreleased ([#330](https://github.com/michaeldcanady/servicenow-sdk-go/issues/330)) ([714b67a](https://github.com/michaeldcanady/servicenow-sdk-go/commit/714b67afd0f3893718fabc63396e7cfe0efd8c70))
* **ci:** use setup-go built-in caching to prevent cache pollution ([#341](https://github.com/michaeldcanady/servicenow-sdk-go/issues/341)) ([f837777](https://github.com/michaeldcanady/servicenow-sdk-go/commit/f837777f003fda850b102ba312cdba1c6e6e27a4)), closes [#339](https://github.com/michaeldcanady/servicenow-sdk-go/issues/339)
* correct key capitalization ([974c125](https://github.com/michaeldcanady/servicenow-sdk-go/commit/974c125dab83996fa2f96aff4d520ae0de1c9104))
* merge error ([13d1eb5](https://github.com/michaeldcanady/servicenow-sdk-go/commit/13d1eb5c41c7afc2f4c10f16861007f7a57affd9))
* not include component name in tag ([83caa5e](https://github.com/michaeldcanady/servicenow-sdk-go/commit/83caa5e08fb06574e44793871a8c539b87357334))
* **paging:** ensure Link headers are correctly captured and parsed ([58e43b6](https://github.com/michaeldcanady/servicenow-sdk-go/commit/58e43b6706eef8b27df99b9891d447fe042f6ef3))
* **REFRESH_TOKEN:** Fixed a trailing whitespace in the refresh_token parameter ([#303](https://github.com/michaeldcanady/servicenow-sdk-go/issues/303)) ([a88a4b2](https://github.com/michaeldcanady/servicenow-sdk-go/commit/a88a4b27013a7788b4fa0d0023ade64552b67bee))
* remove "v" prefix ([8a5b940](https://github.com/michaeldcanady/servicenow-sdk-go/commit/8a5b940bad3d61522ce041dba8ecbcc8c6c112af))
* resolve integration test failures and SDK URL bugs ([#345](https://github.com/michaeldcanady/servicenow-sdk-go/issues/345)) ([d8c4afc](https://github.com/michaeldcanady/servicenow-sdk-go/commit/d8c4afc62eacceaa7dd1bed70d1f55e36586a015))
* strip unneeded comment ([71af4f9](https://github.com/michaeldcanady/servicenow-sdk-go/commit/71af4f9bcb98d088c3daf6f88a13a5f5b11faf3d))
* unable to get response from new table api implementation ([ce520ba](https://github.com/michaeldcanady/servicenow-sdk-go/commit/ce520ba12946040e557f669abef773e454590842))

## [1.10.0](https://github.com/michaeldcanady/servicenow-sdk-go/compare/github.com/michaeldcanady/servicenow-sdk-go-v1.9.0...github.com/michaeldcanady/servicenow-sdk-go-v1.10.0) (2026-03-06)


### Features

* add bug-reporter gemini skill ([#340](https://github.com/michaeldcanady/servicenow-sdk-go/issues/340)) ([f91fa1b](https://github.com/michaeldcanady/servicenow-sdk-go/commit/f91fa1b935040502a0bb29cc0c4d33dcd7a06bfe))
* add support for ropc credentials ([01d7365](https://github.com/michaeldcanady/servicenow-sdk-go/commit/01d7365e142c1ec48606876de98f9890d79bab73))
* **internal:** add request configuration helper and update page iterator ([7fcbf53](https://github.com/michaeldcanady/servicenow-sdk-go/commit/7fcbf534367210e8850c6e5b52fa77d1c8922836))
* make data accessible ([7471d49](https://github.com/michaeldcanady/servicenow-sdk-go/commit/7471d4955a98fa838443accaa672dbb127faa380))
* Reimplementing ropc credentials ([#334](https://github.com/michaeldcanady/servicenow-sdk-go/issues/334)) ([ecc16ae](https://github.com/michaeldcanady/servicenow-sdk-go/commit/ecc16aed98a9c622d48fdc007d135bd4c61382ab))


### Bug Fixes

* add serialization to table record ([aeeb45e](https://github.com/michaeldcanady/servicenow-sdk-go/commit/aeeb45e4d0f9cd591ec9d3909785acb6b8b09371))
* allow serialization ([29be0fd](https://github.com/michaeldcanady/servicenow-sdk-go/commit/29be0fd1503f2936d3283d03963fd879bd6535bd))
* **ci:** use setup-go built-in caching to prevent cache pollution ([#341](https://github.com/michaeldcanady/servicenow-sdk-go/issues/341)) ([f837777](https://github.com/michaeldcanady/servicenow-sdk-go/commit/f837777f003fda850b102ba312cdba1c6e6e27a4)), closes [#339](https://github.com/michaeldcanady/servicenow-sdk-go/issues/339)
* correct key capitalization ([974c125](https://github.com/michaeldcanady/servicenow-sdk-go/commit/974c125dab83996fa2f96aff4d520ae0de1c9104))
* **paging:** ensure Link headers are correctly captured and parsed ([58e43b6](https://github.com/michaeldcanady/servicenow-sdk-go/commit/58e43b6706eef8b27df99b9891d447fe042f6ef3))
* resolve integration test failures and SDK URL bugs ([#345](https://github.com/michaeldcanady/servicenow-sdk-go/issues/345)) ([d8c4afc](https://github.com/michaeldcanady/servicenow-sdk-go/commit/d8c4afc62eacceaa7dd1bed70d1f55e36586a015))
* unable to get response from new table api implementation ([ce520ba](https://github.com/michaeldcanady/servicenow-sdk-go/commit/ce520ba12946040e557f669abef773e454590842))

## [1.9.0](https://github.com/michaeldcanady/servicenow-sdk-go/compare/github.com/michaeldcanady/servicenow-sdk-go-v1.8.0...github.com/michaeldcanady/servicenow-sdk-go-v1.9.0) (2026-03-02)


### Features

* added Last method ([c7b440f](https://github.com/michaeldcanady/servicenow-sdk-go/commit/c7b440f9ccaa2638a5151e168b37ff58258757f9))
* added Last method ([f01d7a0](https://github.com/michaeldcanady/servicenow-sdk-go/commit/f01d7a0c5dd06feac1c21aa3b0c03edde36acd84))
* added Last method ([846052c](https://github.com/michaeldcanady/servicenow-sdk-go/commit/846052ce27f7717811d514d993bcec9d55a28704))
* added Last method ([28f6ccc](https://github.com/michaeldcanady/servicenow-sdk-go/commit/28f6ccc35444711084b4131864b64c05ca9e16f3))


### Bug Fixes

* add missing release-please version tags ([#328](https://github.com/michaeldcanady/servicenow-sdk-go/issues/328)) ([fa83ff1](https://github.com/michaeldcanady/servicenow-sdk-go/commit/fa83ff17cc09bb7a93091f8338d3f1693cef06f5))
* bad ref name ([#331](https://github.com/michaeldcanady/servicenow-sdk-go/issues/331)) ([d64cb4a](https://github.com/michaeldcanady/servicenow-sdk-go/commit/d64cb4ae95f6d8b0ec966e7798f88bc670027009))
* checkout code prior to updating unreleased ([#330](https://github.com/michaeldcanady/servicenow-sdk-go/issues/330)) ([714b67a](https://github.com/michaeldcanady/servicenow-sdk-go/commit/714b67afd0f3893718fabc63396e7cfe0efd8c70))
* **REFRESH_TOKEN:** Fixed a trailing whitespace in the refresh_token parameter ([#303](https://github.com/michaeldcanady/servicenow-sdk-go/issues/303)) ([a88a4b2](https://github.com/michaeldcanady/servicenow-sdk-go/commit/a88a4b27013a7788b4fa0d0023ade64552b67bee))
* remove "v" prefix ([8a5b940](https://github.com/michaeldcanady/servicenow-sdk-go/commit/8a5b940bad3d61522ce041dba8ecbcc8c6c112af))
* strip unneeded comment ([71af4f9](https://github.com/michaeldcanady/servicenow-sdk-go/commit/71af4f9bcb98d088c3daf6f88a13a5f5b11faf3d))
