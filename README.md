# go-simple-telementry

---
## 설명
- 해당 프로젝트는 trace와 metric 관련 오픈소스인  Opentelemetry에 대해 사용하는 방법에 대해 연습 및 실험을 하는 레포지토리입니다.
## Opentelemetry에 대하여
### Opentelemetry란?
Opentelemetry는 log,metric,trace와 같은 데이터들을 생성 및 관리를 위한 APIs,SDKs, tooling 등의 묶음 기술입니다.<br/>
이 프로젝트는 이 프로젝트에서는 선택한 백엔드로 텔레메트리 데이터를 전송하도록 구성할 수 있는 벤더에 의존하지 않는 구현을 제공하며 jaeger, prometheus와 같은 측정 관련 오픈소스 프로그램을 지원해줍니다.
### Opentracing 과 Opencensus
OpenCensus 는 응용 프로그램을 계측하고, 통계(메트릭)를 수집하고, 지원되는 백엔드로 데이터를 내보내기 위한 언어별 라이브러리 모음입니다.<br/> 
OpenTracing은 추적을 위한 표준화된 API이며 개발자가 분산 추적을 위해 자체 서비스 또는 라이브러리를 계측하는 데 사용할 수 있는 사양을 제공합니다.<br/>
### Opentelemtry Components
Opentelemtry는 현재 다음과 같은 다양한 메인 컴포넌트들로 구성되어 있습니다.
- Cross-language specification
- Tools to collect, transform, and export telemetry data
- Per-language SDKs
- Automatic instrumentation and contrib packages
#### Specification
#### Collector
#### Language SDKs
#### Automatic Instrumentation 

