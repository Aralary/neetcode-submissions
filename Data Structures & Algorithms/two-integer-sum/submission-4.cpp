class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, vector<int>> stat;
        for (int i = 0 ; i < nums.size() ; ++i){
            stat[nums[i]].push_back(i);
        }
        for (const auto p : stat){
            auto delta = target - p.first;
            if(stat.count(delta)){
                if (delta == p.first) {
                    if (p.second.size() > 1) {
                        return {stat[delta][0], stat[delta][1]};
                    }
                } else {
                    return {min(p.second[0], stat[delta][0]),max(p.second[0], stat[delta][0])};
                }
            }
        }
    }
};
