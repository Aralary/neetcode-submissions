class Solution {
public:
    bool isAnagram(string s, string t) {
        unordered_map<char, int> data;
        if (s.size() != t.size()) { return false; }
        for (const auto& sep : s){
            data[sep]++;
        }
        for(const auto& sep : t){
            if (!data.count(sep)) {
                return false;
            } else {
                data[sep]--;
                if (data[sep] == 0) {
                    data.erase(sep);
                }
            }
        }
        return data.empty();
    }
};
