{{/*
Expand the name of the chart.
*/}}
{{- define "tempo.name" -}}
{{- $default := "tempo" }}
{{- coalesce .Values.nameOverride .Values.tempo.nameOverride $default | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "tempo.labels" -}}
helm.sh/chart: {{ include "tempo.chart" .ctx }}
{{ include "tempo.selectorLabels" . }}
{{- if .memberlist }}
app.kubernetes.io/part-of: memberlist
{{- end }}
{{- if .ctx.Chart.AppVersion }}
app.kubernetes.io/version: {{ .ctx.Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .ctx.Release.Service }}
giantswarm.io/service-type: "managed"
{{- if and .ctx.Chart .ctx.Chart.Annotations }}
application.giantswarm.io/team: {{ index .ctx.Chart.Annotations "application.giantswarm.io/team" | default "atlas" | quote }}
{{- else if and .Chart .Chart.Annotations }}
application.giantswarm.io/team: {{ index .Chart.Annotations "application.giantswarm.io/team" | default "atlas" | quote }}
{{- else }}
application.giantswarm.io/team: atlas
{{- end }}
{{- end }}
