local promqlouts = import './promqlouts.libsonnet';
{
  new():: {
    out_ports: {
      output: error 'Port output is missing',
    },
  },
  outPorts:: promqlouts,
  withEvaluationInterval(evaluation_interval):: {
    evaluation_interval: evaluation_interval,
  },
  withEvaluationIntervalMixin(evaluation_interval):: {
    evaluation_interval+: evaluation_interval,
  },
  withOutPorts(out_ports):: {
    out_ports: out_ports,
  },
  withOutPortsMixin(out_ports):: {
    out_ports+: out_ports,
  },
  withQueryString(query_string):: {
    query_string: query_string,
  },
  withQueryStringMixin(query_string):: {
    query_string+: query_string,
  },
}
