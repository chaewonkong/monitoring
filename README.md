# 🔥 Monitoring

Simple monitoring system with Prometheus and Grafana.

## Architecture

         +-------------------+
         |                   |
         |        WEB        |
         |    Application    |
         |       Server      |
         |     (Go/Echo)     |
         |                   |
         +---------+---------+
                   |
                   | Metrics & Logs
                   |
         +---------v---------+          +-------------------+
         |                   |          |                   |
         |   Prometheus      |<-------->|   Grafana         |
         |                   |          |                   |
         +-------------------+          +-------------------+
