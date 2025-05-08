#include <iostream>
#include <vector>
using namespace std;

// Reuse add and multiply from above

int main() {
    int n;
    cin >> n;
    vector<vector<int>> A(n, vector<int>(n)), B(n, vector<int>(n)), C(n, vector<int>(n, 0));

    for (int i = 0; i < n; ++i)
        for (int j = 0; j < n; ++j)
            cin >> A[i][j];

    for (int i = 0; i < n; ++i)
        for (int j = 0; j < n; ++j)
            cin >> B[i][j];

    for (int i = 0; i < n; ++i)
        for (int j = 0; j < n; ++j)
            for (int k = 0; k < n; ++k)
                C[i][j] = add(C[i][j], multiply(A[i][k], B[k][j]));

    for (auto& row : C) {
        for (int val : row)
            cout << val << " ";
        cout << "\n";
    }

    return 0;
}
