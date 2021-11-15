package {{ Header.Package  }}

import (
{% for pkg in Header.Imports %}
    "{{ pkg }}"{% endfor %}
)

func TestInsert(t *testing.T) {
    var err error
    mapper := new(MakeMapper)
    for i:=0;i<10;i++{
    {% for testCase in Cases %}
    {{ testCase.Code }}
    {% endfor %}
    }
}

