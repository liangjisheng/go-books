#include "trie.hpp"
#include <iostream>
#include <vector>

using std::vector;
using std::string;

int main() {
    Trie trie;

	trie.insert("apple");
	std::cout << trie.search("apple") << std::endl;    // 返回 true
	std::cout << trie.search("app") << std::endl;      // 返回 false
	std::cout << trie.prefixMatched("app") << std::endl; // 返回 true
	trie.insert("app");
	std::cout << trie.search("app") << std::endl; // 返回 true
	trie.insert("abc");
	
	vector<string> words = trie.getAllWords();
	for (size_t i = 0; i < words.size(); i++) {
		std::cout << words[i] << " ";
	}
	std::cout << std::endl;
	
	trie.deleteWord("abc");
	words = trie.getAllWords();
	for (size_t i = 0; i < words.size(); i++) {
		std::cout << words[i] << " ";
	}
	std::cout << std::endl;
	
	trie.deleteWord("app");
	words = trie.getAllWords();
	for (size_t i = 0; i < words.size(); i++) {
		std::cout << words[i] << " ";
	}
	std::cout << std::endl;

    return 0;
}
