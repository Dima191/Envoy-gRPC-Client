# xds

## Описание

`xds` — это gRPC сервер, предназначенный для взаимодействия с Envoy, обеспечивающий динамическое управление конфигурацией сервисов. Он реализует API XDS (например, LDS, RDS, CDS, EDS), позволяя централизованно управлять маршрутизацией, кластеризацией и другими аспектами работы Envoy.

## Особенности

- **Динамическое обновление конфигурации**: Поддержка API XDS позволяет обновлять конфигурацию Envoy в реальном времени без перезагрузки.
- **Гибкость и масштабируемость**: Легко интегрируется с различными сервисами и масштабируется в зависимости от нагрузки.
- **Incremental Updates**: Delta requests allow for incremental updates to Envoy's configuration, reducing the amount of data sent over the network.
- **Efficiency**: They minimize the impact on the system by updating only the changes, which enhances performance and responsiveness.



## Интеграция с другими серверами
Вместе с xds, в системе работают следующие серверы:

- **route_server**: Сервер для получения эндпоинтов и маршрутов. Он взаимодействует с xds для предоставления актуальной информации о маршрутизации.
- **cert_server**: Сервер для управления сертификатами и безопасным подключением. xds использует его для получения актуальных сертификатов и настроек безопасности.