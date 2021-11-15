package {{ Header.Package  }}

import (
{% for pkg in Header.Imports %}
    "{{ pkg }}"{% endfor %}
)

type MakeMapper struct {
{% for method in Methods %}
    {{ method.Name }} func({% for param in method.In %} {{ param.Name }} {{ param.Type }}, {% endfor %})({% for param in method.Out %} {{ param.Name }} {{ param.Type }}, {% endfor %}){% endfor %}
}