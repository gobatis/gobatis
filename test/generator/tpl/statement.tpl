<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//gobatis.co//DTD Mapper 1.0//EN" "gobatis.co/dtd/mapper.dtd">
<mapper>
{% for statement in Statements %}
<{{statement.Tag}} id="{{ statement.Id }}" {% if statement.ShowParameter %}parameter="{{ RenderParams(statement.Params) }}"{% endif %} {% if statement.ShowResult %}result="{{ RenderParams(statement.Result) }}"{% endif %}>
    {{ statement.Sql }}
</{{statement.Tag}}>
{% endfor %}
</mapper>