# Change Log

## [v0.6.0-rc.0](https://github.com/kubedb/memcached/tree/v0.6.0-rc.0) (2019-08-22)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.5.0...v0.6.0-rc.0)

**Merged pull requests:**

- Improve test: Use installed memcachedversions [\#131](https://github.com/kubedb/memcached/pull/131) ([the-redback](https://github.com/the-redback))
- Update dependencies [\#130](https://github.com/kubedb/memcached/pull/130) ([tamalsaha](https://github.com/tamalsaha))
- Don't set annotation to AppBinding [\#129](https://github.com/kubedb/memcached/pull/129) ([hossainemruz](https://github.com/hossainemruz))
- Set database version in AppBinding [\#128](https://github.com/kubedb/memcached/pull/128) ([hossainemruz](https://github.com/hossainemruz))
- Change package path to kubedb.dev/memcached [\#127](https://github.com/kubedb/memcached/pull/127) ([tamalsaha](https://github.com/tamalsaha))
- Add license header to Makefiles [\#126](https://github.com/kubedb/memcached/pull/126) ([tamalsaha](https://github.com/tamalsaha))
- Add install, uninstall and purge command in Makefile [\#125](https://github.com/kubedb/memcached/pull/125) ([hossainemruz](https://github.com/hossainemruz))
- Add Makefile [\#124](https://github.com/kubedb/memcached/pull/124) ([tamalsaha](https://github.com/tamalsaha))
- Pod Disruption Budget for Memcached [\#123](https://github.com/kubedb/memcached/pull/123) ([iamrz1](https://github.com/iamrz1))
- Handling resource ownership [\#122](https://github.com/kubedb/memcached/pull/122) ([iamrz1](https://github.com/iamrz1))
- Update to k8s 1.14.0 client libraries using go.mod [\#121](https://github.com/kubedb/memcached/pull/121) ([tamalsaha](https://github.com/tamalsaha))

## [0.5.0](https://github.com/kubedb/memcached/tree/0.5.0) (2019-05-06)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.4.0...0.5.0)

**Merged pull requests:**

- Revendor dependencies [\#120](https://github.com/kubedb/memcached/pull/120) ([tamalsaha](https://github.com/tamalsaha))
- Fix PSP in Role for kubeDB upgrade [\#119](https://github.com/kubedb/memcached/pull/119) ([iamrz1](https://github.com/iamrz1))
- Modify mutator validator names [\#118](https://github.com/kubedb/memcached/pull/118) ([iamrz1](https://github.com/iamrz1))

## [0.4.0](https://github.com/kubedb/memcached/tree/0.4.0) (2019-03-18)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.3.0...0.4.0)

**Merged pull requests:**

- DB psp in e2e test framework [\#117](https://github.com/kubedb/memcached/pull/117) ([iamrz1](https://github.com/iamrz1))
- Don't inherit app.kubernetes.io labels from CRD into offshoots [\#116](https://github.com/kubedb/memcached/pull/116) ([tamalsaha](https://github.com/tamalsaha))
- Add role label to stats service [\#115](https://github.com/kubedb/memcached/pull/115) ([tamalsaha](https://github.com/tamalsaha))
- PSP support for Memcached [\#114](https://github.com/kubedb/memcached/pull/114) ([iamrz1](https://github.com/iamrz1))
- Update Kubernetes client libraries to 1.13.0 release [\#113](https://github.com/kubedb/memcached/pull/113) ([tamalsaha](https://github.com/tamalsaha))

## [0.3.0](https://github.com/kubedb/memcached/tree/0.3.0) (2019-02-19)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.2.0...0.3.0)

**Merged pull requests:**

- Revendor dependencies [\#112](https://github.com/kubedb/memcached/pull/112) ([tamalsaha](https://github.com/tamalsaha))
- Initial RBAC support: create and use K8s service account for Memcached [\#111](https://github.com/kubedb/memcached/pull/111) ([the-redback](https://github.com/the-redback))
-  Revendor dependencies [\#110](https://github.com/kubedb/memcached/pull/110) ([the-redback](https://github.com/the-redback))
- Revendor dependencies [\#109](https://github.com/kubedb/memcached/pull/109) ([the-redback](https://github.com/the-redback))
- Add certificate health checker [\#108](https://github.com/kubedb/memcached/pull/108) ([tamalsaha](https://github.com/tamalsaha))
- Update E2E test: Env update is not restricted anymore [\#107](https://github.com/kubedb/memcached/pull/107) ([the-redback](https://github.com/the-redback))
- Fix app binding [\#106](https://github.com/kubedb/memcached/pull/106) ([tamalsaha](https://github.com/tamalsaha))

## [0.2.0](https://github.com/kubedb/memcached/tree/0.2.0) (2018-12-17)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.2.0-rc.2...0.2.0)

**Merged pull requests:**

- Reuse event recorder [\#105](https://github.com/kubedb/memcached/pull/105) ([tamalsaha](https://github.com/tamalsaha))
- Revendor dependencies [\#104](https://github.com/kubedb/memcached/pull/104) ([tamalsaha](https://github.com/tamalsaha))

## [0.2.0-rc.2](https://github.com/kubedb/memcached/tree/0.2.0-rc.2) (2018-12-06)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.2.0-rc.1...0.2.0-rc.2)

**Merged pull requests:**

- Ignore mutation of fields to default values during update [\#102](https://github.com/kubedb/memcached/pull/102) ([tamalsaha](https://github.com/tamalsaha))
- Support configuration options for exporter sidecar [\#101](https://github.com/kubedb/memcached/pull/101) ([tamalsaha](https://github.com/tamalsaha))
- Use flags.DumpAll [\#100](https://github.com/kubedb/memcached/pull/100) ([tamalsaha](https://github.com/tamalsaha))

## [0.2.0-rc.1](https://github.com/kubedb/memcached/tree/0.2.0-rc.1) (2018-12-02)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.2.0-rc.0...0.2.0-rc.1)

**Merged pull requests:**

- Apply cleanup [\#99](https://github.com/kubedb/memcached/pull/99) ([tamalsaha](https://github.com/tamalsaha))
- Set periodic analytics [\#98](https://github.com/kubedb/memcached/pull/98) ([tamalsaha](https://github.com/tamalsaha))
- Introduce AppBinding support [\#97](https://github.com/kubedb/memcached/pull/97) ([the-redback](https://github.com/the-redback))
- Fix analytics [\#96](https://github.com/kubedb/memcached/pull/96) ([the-redback](https://github.com/the-redback))
- Fix operator startup in minikube [\#95](https://github.com/kubedb/memcached/pull/95) ([the-redback](https://github.com/the-redback))
- Add CRDS without observation when operator starts [\#94](https://github.com/kubedb/memcached/pull/94) ([the-redback](https://github.com/the-redback))

## [0.2.0-rc.0](https://github.com/kubedb/memcached/tree/0.2.0-rc.0) (2018-10-15)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.2.0-beta.1...0.2.0-rc.0)

**Merged pull requests:**

- Support providing resources for monitoring container [\#93](https://github.com/kubedb/memcached/pull/93) ([tamalsaha](https://github.com/tamalsaha))
- Update kubernetes client libraries to 1.12.0 [\#92](https://github.com/kubedb/memcached/pull/92) ([tamalsaha](https://github.com/tamalsaha))
- Add validation webhook xray [\#91](https://github.com/kubedb/memcached/pull/91) ([tamalsaha](https://github.com/tamalsaha))
- Various Fixes [\#90](https://github.com/kubedb/memcached/pull/90) ([hossainemruz](https://github.com/hossainemruz))
- Merge ports from service template [\#88](https://github.com/kubedb/memcached/pull/88) ([tamalsaha](https://github.com/tamalsaha))
- Replace doNotPause with TerminationPolicy = DoNotTerminate [\#87](https://github.com/kubedb/memcached/pull/87) ([tamalsaha](https://github.com/tamalsaha))
- Pass resources to NamespaceValidator [\#86](https://github.com/kubedb/memcached/pull/86) ([tamalsaha](https://github.com/tamalsaha))
- Various fixes [\#85](https://github.com/kubedb/memcached/pull/85) ([tamalsaha](https://github.com/tamalsaha))
- Support Livecycle hook and container probes [\#84](https://github.com/kubedb/memcached/pull/84) ([tamalsaha](https://github.com/tamalsaha))
- Check if Kubernetes version is supported before running operator [\#83](https://github.com/kubedb/memcached/pull/83) ([tamalsaha](https://github.com/tamalsaha))
- Update package alias [\#82](https://github.com/kubedb/memcached/pull/82) ([tamalsaha](https://github.com/tamalsaha))

## [0.2.0-beta.1](https://github.com/kubedb/memcached/tree/0.2.0-beta.1) (2018-09-30)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.2.0-beta.0...0.2.0-beta.1)

**Merged pull requests:**

- Revendor api [\#81](https://github.com/kubedb/memcached/pull/81) ([tamalsaha](https://github.com/tamalsaha))
- Fix tests [\#80](https://github.com/kubedb/memcached/pull/80) ([tamalsaha](https://github.com/tamalsaha))
- Revendor api for catalog apigroup [\#79](https://github.com/kubedb/memcached/pull/79) ([tamalsaha](https://github.com/tamalsaha))
- Use --pull flag with docker build \(\#20\) [\#78](https://github.com/kubedb/memcached/pull/78) ([tamalsaha](https://github.com/tamalsaha))

## [0.2.0-beta.0](https://github.com/kubedb/memcached/tree/0.2.0-beta.0) (2018-09-20)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.1.0...0.2.0-beta.0)

**Merged pull requests:**

- Show deprecated column for memcachedversions [\#77](https://github.com/kubedb/memcached/pull/77) ([hossainemruz](https://github.com/hossainemruz))
- Support Termination Policy & Stop working for deprecated \*Versions [\#76](https://github.com/kubedb/memcached/pull/76) ([the-redback](https://github.com/the-redback))
- Revendor k8s.io/apiserver [\#75](https://github.com/kubedb/memcached/pull/75) ([tamalsaha](https://github.com/tamalsaha))
- Revendor kubernetes-1.11.3 [\#74](https://github.com/kubedb/memcached/pull/74) ([tamalsaha](https://github.com/tamalsaha))
- Support UpdateStrategy [\#73](https://github.com/kubedb/memcached/pull/73) ([tamalsaha](https://github.com/tamalsaha))
- Add TerminationPolicy for databases [\#72](https://github.com/kubedb/memcached/pull/72) ([tamalsaha](https://github.com/tamalsaha))
- Revendor api [\#71](https://github.com/kubedb/memcached/pull/71) ([tamalsaha](https://github.com/tamalsaha))
- Use IntHash as status.observedGeneration [\#70](https://github.com/kubedb/memcached/pull/70) ([tamalsaha](https://github.com/tamalsaha))
- fix github status [\#69](https://github.com/kubedb/memcached/pull/69) ([tahsinrahman](https://github.com/tahsinrahman))
- update pipeline [\#68](https://github.com/kubedb/memcached/pull/68) ([tahsinrahman](https://github.com/tahsinrahman))
- Fix E2E test for minikube [\#67](https://github.com/kubedb/memcached/pull/67) ([the-redback](https://github.com/the-redback))
- update pipeline [\#66](https://github.com/kubedb/memcached/pull/66) ([tahsinrahman](https://github.com/tahsinrahman))
- Migrate memcached [\#65](https://github.com/kubedb/memcached/pull/65) ([tamalsaha](https://github.com/tamalsaha))
- Use official exporter image [\#64](https://github.com/kubedb/memcached/pull/64) ([the-redback](https://github.com/the-redback))
- Update status.ObservedGeneration for failure phase [\#63](https://github.com/kubedb/memcached/pull/63) ([the-redback](https://github.com/the-redback))
- Keep track of ObservedGenerationHash [\#62](https://github.com/kubedb/memcached/pull/62) ([tamalsaha](https://github.com/tamalsaha))
- Use NewObservableHandler [\#61](https://github.com/kubedb/memcached/pull/61) ([tamalsaha](https://github.com/tamalsaha))
- Fix uninstall for concourse [\#60](https://github.com/kubedb/memcached/pull/60) ([tahsinrahman](https://github.com/tahsinrahman))
- Revise immutable spec fields [\#59](https://github.com/kubedb/memcached/pull/59) ([tamalsaha](https://github.com/tamalsaha))
- Support passing args via PodTemplate [\#58](https://github.com/kubedb/memcached/pull/58) ([tamalsaha](https://github.com/tamalsaha))
- Revendor api [\#57](https://github.com/kubedb/memcached/pull/57) ([tamalsaha](https://github.com/tamalsaha))
- Add support for running tests on cncf cluster [\#56](https://github.com/kubedb/memcached/pull/56) ([tahsinrahman](https://github.com/tahsinrahman))
- Keep track of observedGeneration in status [\#55](https://github.com/kubedb/memcached/pull/55) ([tamalsaha](https://github.com/tamalsaha))
- Separate StatsService for monitoring [\#54](https://github.com/kubedb/memcached/pull/54) ([shudipta](https://github.com/shudipta))
-  Use MemcachedVersion for Memcached images [\#53](https://github.com/kubedb/memcached/pull/53) ([the-redback](https://github.com/the-redback))
- Use updated crd spec [\#52](https://github.com/kubedb/memcached/pull/52) ([tamalsaha](https://github.com/tamalsaha))
- Rename OffshootLabels to OffshootSelectors [\#51](https://github.com/kubedb/memcached/pull/51) ([tamalsaha](https://github.com/tamalsaha))
- Revendor api [\#50](https://github.com/kubedb/memcached/pull/50) ([tamalsaha](https://github.com/tamalsaha))
- Use kmodules monitoring and objectstore api [\#49](https://github.com/kubedb/memcached/pull/49) ([tamalsaha](https://github.com/tamalsaha))
- Refactor concourse scripts [\#48](https://github.com/kubedb/memcached/pull/48) ([tahsinrahman](https://github.com/tahsinrahman))
- Fix command `./hack/make.py test e2e` [\#47](https://github.com/kubedb/memcached/pull/47) ([the-redback](https://github.com/the-redback))
- Support custom configuration [\#46](https://github.com/kubedb/memcached/pull/46) ([hossainemruz](https://github.com/hossainemruz))
- Set generated binary name to mc-operator [\#45](https://github.com/kubedb/memcached/pull/45) ([tamalsaha](https://github.com/tamalsaha))
- Don't add admission/v1beta1 group as a prioritized version [\#44](https://github.com/kubedb/memcached/pull/44) ([tamalsaha](https://github.com/tamalsaha))
- Enable status subresource for crds [\#43](https://github.com/kubedb/memcached/pull/43) ([tamalsaha](https://github.com/tamalsaha))
- Update client-go to v8.0.0 [\#42](https://github.com/kubedb/memcached/pull/42) ([tamalsaha](https://github.com/tamalsaha))
- Format shell script [\#41](https://github.com/kubedb/memcached/pull/41) ([tamalsaha](https://github.com/tamalsaha))
- Support ENV variables in CRDs [\#39](https://github.com/kubedb/memcached/pull/39) ([hossainemruz](https://github.com/hossainemruz))

## [0.1.0](https://github.com/kubedb/memcached/tree/0.1.0) (2018-06-12)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.1.0-rc.0...0.1.0)

**Merged pull requests:**

-  Fixed missing error return [\#38](https://github.com/kubedb/memcached/pull/38) ([the-redback](https://github.com/the-redback))
- Revendor dependencies [\#37](https://github.com/kubedb/memcached/pull/37) ([tamalsaha](https://github.com/tamalsaha))
- Add changelog [\#36](https://github.com/kubedb/memcached/pull/36) ([tamalsaha](https://github.com/tamalsaha))

## [0.1.0-rc.0](https://github.com/kubedb/memcached/tree/0.1.0-rc.0) (2018-05-28)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.1.0-beta.2...0.1.0-rc.0)

**Merged pull requests:**

- Fixed kubeconfig plugin for Cloud Providers [\#35](https://github.com/kubedb/memcached/pull/35) ([the-redback](https://github.com/the-redback))
- Concourse [\#34](https://github.com/kubedb/memcached/pull/34) ([tahsinrahman](https://github.com/tahsinrahman))
- Refactored E2E testing to support E2E testing with admission webhook in cloud [\#33](https://github.com/kubedb/memcached/pull/33) ([the-redback](https://github.com/the-redback))
- Skip delete requests for empty resources [\#32](https://github.com/kubedb/memcached/pull/32) ([the-redback](https://github.com/the-redback))
- Don't panic if admission options is nil [\#31](https://github.com/kubedb/memcached/pull/31) ([tamalsaha](https://github.com/tamalsaha))
- Disable admission controllers for webhook server [\#30](https://github.com/kubedb/memcached/pull/30) ([tamalsaha](https://github.com/tamalsaha))
- Separate ApiGroup for Mutating and Validating webhook [\#29](https://github.com/kubedb/memcached/pull/29) ([the-redback](https://github.com/the-redback))
- Update client-go to 7.0.0 [\#28](https://github.com/kubedb/memcached/pull/28) ([tamalsaha](https://github.com/tamalsaha))
-  Bundle admission webhook server and Used SharedInformer [\#27](https://github.com/kubedb/memcached/pull/27) ([the-redback](https://github.com/the-redback))
-  Moved admission webhook packages into memcached repo [\#26](https://github.com/kubedb/memcached/pull/26) ([the-redback](https://github.com/the-redback))
- Add travis yaml [\#25](https://github.com/kubedb/memcached/pull/25) ([tahsinrahman](https://github.com/tahsinrahman))

## [0.1.0-beta.2](https://github.com/kubedb/memcached/tree/0.1.0-beta.2) (2018-02-27)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.1.0-beta.1...0.1.0-beta.2)

**Merged pull requests:**

-  Migrating to apps/v1 [\#23](https://github.com/kubedb/memcached/pull/23) ([the-redback](https://github.com/the-redback))
- fix for pointer type & update validation [\#22](https://github.com/kubedb/memcached/pull/22) ([aerokite](https://github.com/aerokite))
- Fix dormantDB matching: pass same type to Equal method [\#20](https://github.com/kubedb/memcached/pull/20) ([the-redback](https://github.com/the-redback))
- Use official code generator scripts [\#19](https://github.com/kubedb/memcached/pull/19) ([tamalsaha](https://github.com/tamalsaha))
- Fixed dormantdb matching & Raised throttling time & Fixed Memcached version checking [\#18](https://github.com/kubedb/memcached/pull/18) ([the-redback](https://github.com/the-redback))

## [0.1.0-beta.1](https://github.com/kubedb/memcached/tree/0.1.0-beta.1) (2018-01-29)
[Full Changelog](https://github.com/kubedb/memcached/compare/0.1.0-beta.0...0.1.0-beta.1)

**Merged pull requests:**

- Update dependencies to client-go 6.0.0 [\#17](https://github.com/kubedb/memcached/pull/17) ([tamalsaha](https://github.com/tamalsaha))
- Memcached 1.5 image added [\#16](https://github.com/kubedb/memcached/pull/16) ([the-redback](https://github.com/the-redback))
- Fixed logger and analytics [\#15](https://github.com/kubedb/memcached/pull/15) ([the-redback](https://github.com/the-redback))
- Refactored Monitoing Management [\#14](https://github.com/kubedb/memcached/pull/14) ([the-redback](https://github.com/the-redback))
- Update docker build scripts [\#13](https://github.com/kubedb/memcached/pull/13) ([tamalsaha](https://github.com/tamalsaha))

## [0.1.0-beta.0](https://github.com/kubedb/memcached/tree/0.1.0-beta.0) (2018-01-07)
**Merged pull requests:**

- Fix Analytics and Rbac [\#12](https://github.com/kubedb/memcached/pull/12) ([the-redback](https://github.com/the-redback))
- Add Docker-registry and Fixed resume Dormant DB [\#10](https://github.com/kubedb/memcached/pull/10) ([the-redback](https://github.com/the-redback))
- Check validation first [\#8](https://github.com/kubedb/memcached/pull/8) ([tamalsaha](https://github.com/tamalsaha))
- Set client id for analytics [\#7](https://github.com/kubedb/memcached/pull/7) ([tamalsaha](https://github.com/tamalsaha))
- Add Workqueue to Memcached-watcher [\#6](https://github.com/kubedb/memcached/pull/6) ([the-redback](https://github.com/the-redback))
- Fix CRD registration procedure [\#5](https://github.com/kubedb/memcached/pull/5) ([the-redback](https://github.com/the-redback))
- Update pkg paths to kubedb org [\#4](https://github.com/kubedb/memcached/pull/4) ([tamalsaha](https://github.com/tamalsaha))
- Assign default Prometheus Monitoring Port [\#3](https://github.com/kubedb/memcached/pull/3) ([the-redback](https://github.com/the-redback))
- Initial implementation [\#2](https://github.com/kubedb/memcached/pull/2) ([the-redback](https://github.com/the-redback))



\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*