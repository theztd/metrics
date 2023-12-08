# Server metrics exporter   

My custom additional netrics exporter for linux servers.


## Use


### Simple ussage (cron and textfile exporter)

```cron
*/5 * * * *  /usr/local/bin/metrics > /var/metrics/my_metrics.prom
```

**For more options run binary with -h flag.**

