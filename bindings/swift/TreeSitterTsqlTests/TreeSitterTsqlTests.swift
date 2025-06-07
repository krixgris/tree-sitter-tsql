import XCTest
import SwiftTreeSitter
import TreeSitterTsql

final class TreeSitterTsqlTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_tsql())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Transact-SQL grammar")
    }
}
