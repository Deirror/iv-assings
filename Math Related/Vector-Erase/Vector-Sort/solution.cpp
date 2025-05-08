#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */   
    int n ;
    cin >> n;
    
    vector<int> vec(n);
    
    for(int i = 0; i < n; i++) {
        cin >> vec[i];
    }
    
    sort(vec.begin(), vec.end());

    for(auto num : vec) {
        cout << num << ' ';
    }
    
    return 0;
}
