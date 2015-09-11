require.config({
    baseUrl: "/static/",
    paths: {
        jquery: "js/lib/jquery-2.1.4.min",
        semantic: "semantic/js/semantic.min",
        validate: "js/lib/jquery.validate-1.13.1.min",
        marked: "editormd/marked.min",
        prettify: "editormd/prettify.min",
        raphael: "editormd/raphael.min",
        underscore: "editormd/underscore.min",
        flowchart: "editormd/flowchart.min",
        jqueryflowchart: "editormd/jquery.flowchart.min",
        sequenceDiagram: "editormd/sequence-diagram.min",
        katex: "editormd/katex.min",
        editormd: "editormd/js/editormd.amd.min" // Using Editor.md amd version for Require.js
    }
});