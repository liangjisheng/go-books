#include <string>
#include <vector>
#include <string.h>

using std::string;
using std::vector;

class Trie {
private:
    bool isEnd;
    Trie* next[26];

    //返回与前缀prefix匹配的最后一个结点的地址
    Trie* __prefix(const string& prefix) {
        Trie* node = this;
        for (char c : prefix) {
            node = node->next[c - 'a'];
            if (node == NULL) {
                return NULL;
            }
        }
        return node;
    }

    //获取以root为起始结点的所有单词
    void __getWords(Trie* root, string& word, vector<string>& allWords) {
        if (root == NULL) {
            return;
        }
        if (root->isEnd) {
            allWords.push_back(word);
        }
        for (int i = 0; i < 26; i++) {
            word.push_back(static_cast<char>('a'+i));
            __getWords(root->next[i], word, allWords);
            word.pop_back();
        }
    }

    //删除一个单词
    void __deleteWord(Trie*& node, const string& word, int d) {
        if (d == word.length()) {
            node->isEnd = false;
        } else {
            __deleteWord(node->next[word[d]-'a'], word, d+1);
        }
        if (node->isEnd) {
            return;
        }
        for (Trie* item : node->next) {
            if (item != NULL) {
                return;
            }
        }
        delete node;
        node = NULL;
    }

public:
    //构造函数
    Trie() {
        isEnd = false;
        memset(next, 0, sizeof(next));
    }

    //插入函数
    void insert(const string& word) {
        Trie* node = this;
        for (char c : word) {
            if (node->next[c-'a'] == NULL) {
                node->next[c-'a'] = new Trie();
            }
            node = node->next[c-'a'];
        }
        node->isEnd = true;
    }

    //查询函数
    bool search(const string& word) {
        Trie* node = __prefix(word);
        if (node == NULL) {
            return false;
        }
        return node->isEnd;
    }

    //前缀匹配函数
    bool prefixMatched(const string& prefix) {
        Trie* node = this;
        for (char c : prefix) {
            node = node->next[c - 'a'];
            if (node == NULL) {
                return false;
            }
        }
        return true;
    }

    //获取以prefix为前缀的所有单词
    vector<string> getAllWordsOfPrefix(const string& prefix) {
        vector<string> words;
        string word = prefix;
        Trie* node = __prefix(prefix);
        __getWords(node, word, words);
        return words;
    }

    //获取Trie中所有单词
    vector<string> getAllWords() {
        string word = "";
        return getAllWordsOfPrefix(word);
    }

    //删除一个单词
    void deleteWord(const string& word) {
        if (!search(word)) {
            return;
        }
        Trie* node = this;
        __deleteWord(node, word, 0);
    }
};