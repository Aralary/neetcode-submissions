class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<pair<int,int>> cp_nums(nums.size());
        for (int i = 0 ; i < nums.size() ; ++i){
            cp_nums[i] = make_pair(nums[i], i);
        }
        sort(cp_nums.begin(), cp_nums.end(), [](pair<int,int>a, pair<int,int>b){
            if (a.first == b.first) {
                return a.second < b.second;
            }else {
                return a.first < b.first;
            }
        });
        int left = 0;
        int right = nums.size() - 1;
        while(left < right) {
            if (cp_nums[left].first + cp_nums[right].first > target) {
                right--;
                continue;
            } else if (cp_nums[left].first + cp_nums[right].first < target) {
                left++;
                continue;
            } else {
                if(left != right) {
                    
                    return {int(min(cp_nums[left].second,cp_nums[right].second)), int(max(cp_nums[left].second,cp_nums[right].second))};
                }
            }
        }
    }
};
