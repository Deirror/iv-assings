class Solution {
public:
    int majorityElement(vector<int>& nums) {
        unordered_map<int, int> hist;
        for(auto num : nums) {
            hist[num]++;
        }
        int times = nums.size() / 2;
        int majorityElement = 0;
        for(auto& kvp : hist) {
            if(kvp.second > times) {
                majorityElement = kvp.first;
            }
        }
        return majorityElement;
    }
};
