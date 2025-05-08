class Solution {
public:
    int dfs(vector<vector<int>>& dp, vector<vector<int>>& mat, int i, int j) {
        int n = mat.size(), m = mat[0].size();

        if (i < 0 || j < 0 || i >= n || j >= m) return 1e5;
        if (mat[i][j] == 0) return 0;

        if (dp[i][j] != -1) return dp[i][j];

        dp[i][j] = 1e5;

        int up    = dfs(dp, mat, i - 1, j);
        int down  = dfs(dp, mat, i + 1, j);
        int left  = dfs(dp, mat, i, j - 1);
        int right = dfs(dp, mat, i, j + 1);

        dp[i][j] = min({up, down, left, right}) + 1;
        return dp[i][j];
    }

    vector<vector<int>> updateMatrix(vector<vector<int>>& mat) {
        int n = mat.size(), m = mat[0].size();
        vector<vector<int>> dp(n, vector<int>(m, -1));

        for (int i = 0; i < n; i++)
            for (int j = 0; j < m; j++)
                if (mat[i][j] == 1)
                    dfs(dp, mat, i, j);
                else
                    dp[i][j] = 0;

        return dp;
    }
};
