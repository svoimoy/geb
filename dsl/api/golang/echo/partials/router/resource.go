{{#with . as |R|}}
{{#methods}}
{{#with list}}
// LIST  {{input}}  ->  {{output}}
G.GET("/{{route}}", resources.Handle_LIST_{{R.name}})
{{/with}}

{{#with get   }}
// GET  {{input}}  ->  {{output}}
G.GET("/{{route}}/:id", resources.Handle_GET_{{R.name}})
{{/with}}

{{#with post   }}
//  POST  {{input}}  ->  {{output}}
G.POST("/{{route}}", resources.Handle_POST_{{R.name}})
{{/with}}

{{#with put   }}
//  PUT  {{input}}  ->  {{output}}
G.PUT("/{{route}}", resources.Handle_PUT_{{R.name}})
{{/with}}

{{#with patch }}
//  PATCH  {{input}}  ->  {{output}}
G.PATCH("/{{route}}/:id", resources.Handle_PATCH_{{R.name}})
{{/with}}

{{#with delete}}
//  DELETE  {{input}}  ->  {{output}}
G.DELETE("/{{route}}/:id", resources.Handle_DELETE_{{R.name}})
{{/with}}
{{/methods}}
{{/with}}

