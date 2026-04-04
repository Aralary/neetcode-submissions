class Solution {
public:
    bool hasDuplicate(vector<int>& nums) {
        unordered_map<int,int> stat;
        for (const auto& i : nums){
            stat[i]++;
            if(stat[i] > 1) return true;
        }
        return false;
    }
};
