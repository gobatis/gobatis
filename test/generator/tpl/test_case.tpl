{% autoescape off %}
package {{ Header.Package  }}

import (
{% for pkg in Header.Imports %}
    "{{ pkg }}"{% endfor %}
)

func testScanTypes(t *testing.T, mapper *Mapper,manager *generator.DataManager) {
    {% for testCase in Cases %}{
        {{ testCase.Code }}
    }
    {% endfor %}
}
{% endautoescape %}

