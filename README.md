# ConverterDataService
Тестовый сервис по преобразованию данных из файла YAML (находится в корне проекта) в формат OpenMetrics

## Пример использования 
```
curl -X GET http://localhost:8080/metrics
```

## Возвращаются данные в формате 
```
<metric name>{<label name>=<label value>, ...}
```
