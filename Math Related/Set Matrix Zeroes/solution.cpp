class Solution {
public:
    void fill(vector<vector<int>>& matrix, int i, int j) {
        int n = matrix.size(),
            m = matrix[0].size();

        bool flag = false;
        for (int k = 1;;k++) {
            if (i + k < n) {
                flag = true;
                matrix[i+k][j] = 0;
            }
            if (j + k < m) {
                flag = true;
                matrix[i][j+k] = 0;
            }
            if (i - k >= 0) {
                flag = true;
                matrix[i-k][j] = 0;
            }
            if (j - k >= 0) {
                flag = true;
                matrix[i][j - k] = 0;
            }
            if (!flag) {
                return;
            }
            flag = false;
        }
    }

    void setZeroes(vector<vector<int>>& matrix) {
        set<pair<int, int>> indexes;

        int n = matrix.size(),
            m = matrix[0].size();

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (matrix[i][j] == 0) {
                    indexes.insert({i, j});
                }
            }
        }

        for (auto& index : indexes) {
            fill(matrix, index.first, index.second);
        }
    }
};
