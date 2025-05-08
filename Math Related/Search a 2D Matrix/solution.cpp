class Solution {
public:
    int search(vector<int>& nums, int target) {
        int l = 0, r = nums.size() - 1;
        while (r >= l) {
            int m = l + (r - l) / 2;
            if (target == nums[m]) return m;
            if (target > nums[m]) l = m + 1;
            else r = m - 1;
        }
        return -1;
    }

    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int n = matrix.size();
        if (n == 0) return false;
        int m = matrix[0].size();
        if (m == 0) return false;

        vector<int> fCol(n);
        for (int i = 0; i < n; i++) {
            fCol[i] = matrix[i][0];
        }

        auto ptr = upper_bound(fCol.begin(), fCol.end(), target);
        int row = (ptr == fCol.begin()) ? 0 : (ptr - fCol.begin() - 1);

        return search(matrix[row], target) != -1;
    }
};
