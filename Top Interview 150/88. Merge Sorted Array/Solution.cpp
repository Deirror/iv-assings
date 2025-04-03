class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        vector<int> result(m + n);
        int i = 0, j = 0, k = 0;
        while(i < m && j < n) {
            if(nums1[i] < nums2[j]) {
                result[k++]=nums1[i++];
            } else {
                result[k++]=nums2[j++];
            }
        }

        while(i < m) {
            result[k++]=nums1[i++];
        }
        
        while(j < n) {
            result[k++]=nums2[j++];
        }

        for(int t = 0; t < result.size(); t++) {
            nums1[t] = result[t];
        }
    }
};
