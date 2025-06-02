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
    
    int toRemove = 0;
    cin >> toRemove;
    
    vec.erase(vec.begin() + toRemove - 1);
    
    pair<int, int> range;
    
    cin >> range.first >> range.second;
    
    vec.erase(vec.begin() + range.first - 1, vec.begin() + range.second - 1);
    
    cout << vec.size() << '\n';
    for(auto num : vec) {
        cout << num << ' ';
    }
    
    return 0;
}
