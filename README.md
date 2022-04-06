# go-simple-telemetry

---
## 설명
- 해당 프로젝트는 trace와 metric 관련 오픈소스인  Opentelemetry에 대해 사용하는 방법에 대해 연습 및 실험을 하는 레포지토리입니다.
- .env_sample을 .env로 변경해서 GCP proejct id를 입력해주시면 됩니다.
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
모든 구현에 대한 언어 간 요구사항 및 기대치를 설명되며 다음과 같은 규격을 가지고 있습니다.
- API: 추적, 메트릭 및 로깅 데이터를 생성하고 상호 연관시키기 위한 데이터 유형 및 작업을 정의합니다.
- SDK: API의 언어별 구현을 위한 요구 사항을 정의합니다. 구성, 데이터 처리 및 내보내기 개념도 여기에 정의되어 있습니다.
- 데이터: 원격 측정 백엔드가 지원할 수 있는 Open Telemetry Protocol(OTLP) 및 공급업체에 구애받지 않는 의미 규칙을 정의합니다.
#### Collector
Open Telemetry Collector는 원격 측정 데이터를 수신, 처리 및 내보낼 수 있는 벤더에 구애받지 않는 프록시입니다. 여러 가지 형식(예: OTLP, jaeger, prometheus 등 다양한 툴) 원격 측정 데이터를 수신하고 하나 이상의 백엔드로 데이터를 전송할 수 있습니다. 또한 원격 측정 데이터를 내보내기 전에 처리하고 필터링할 수도 있습니다. 수집기 기여 패키지는 더 많은 데이터 형식과 공급업체 백엔드를 지원합니다.
#### Language SDKs
OpenTelemetry에는 OpenTelemetry API를 사용하여 선택한 언어로 원격 측정 데이터를 생성하고 해당 데이터를 기본 백엔드로 내보낼 수 있는 언어 SDK도 있습니다. 또한 이러한 SDK를 사용하면 응용프로그램의 수동 계측기에 연결하는 데 사용할 수 있는 공통 라이브러리 및 프레임워크에 대한 자동 계측기를 통합할 수 있습니다. 공급업체들은 종종 그들의 백엔드로 내보내기를 더 쉽게 하기 위해 언어 SDK를 배포한다.
#### Automatic Instrumentation 
Open Telemetry는 널리 사용되는 라이브러리와 지원되는 언어에 대한 프레임워크에서 관련 원격 측정 데이터를 생성하는 광범위한 구성 요소를 지원합니다. 예를 들어 HTTP 라이브러리의 인바운드 및 아웃바운드 HTTP 요청은 해당 요청에 대한 데이터를 생성합니다. 자동 계측기를 사용하는 것은 언어마다 다를 수 있으며, 응용 프로그램과 함께 로드하는 구성 요소를 선호하거나 사용해야 하며, 다른 경우에는 코드베이스에서 패키지를 명시적으로 가져오는 것을 선호할 수 있습니다.

