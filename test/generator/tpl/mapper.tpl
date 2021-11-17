{% autoescape off %}
package {{ Header.Package  }}

import (
{% for pkg in Header.Imports %}
    "{{ pkg }}"{% endfor %}
)

type MakeMapper struct {
{% for method in Methods %}
    {{ method.Name }} func({{ RenderGoParams(method.In) }})({{ RenderGoParams(method.Out) }}){% endfor %}
}
{% endautoescape %}