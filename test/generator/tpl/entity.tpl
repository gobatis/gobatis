{% autoescape off %}
package {{ Header.Package  }}

import (
{% for pkg in Header.Imports %}
    "{{ pkg }}"
{% endfor %}
)

{% for entity in Entities %}
type {{ entity.Name }} struct{ {% for param in entity.Params %}
        {{ param.Name }} {{ param.Type }} `sql:"{{ param.Tag }}"`{% endfor %}
}
{% endfor %}
{% endautoescape %}