class Solution {
public:
    int lengthOfLongestSubstring(string str) {
        unordered_set<char> s;
        int maxLen=0;
        int leftmost=0;
        for(int i = 0; i < str.size(); i++) {
            while(s.find(str[i])!=s.end()) {
                s.erase(str[leftmost++]);
            }
            s.insert(str[i]);
            maxLen=max(maxLen, (int)s.size());
        }
        return maxLen;
    }
};
