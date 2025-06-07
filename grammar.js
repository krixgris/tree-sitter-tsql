/**
 * @file Tree-sitter parser for MSSQL T-SQL inspired by the really powerful alround SQL parser by DerekStride (https://github.com/DerekStride/tree-sitter-sql)
 * @author krixgris
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "tsql",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
