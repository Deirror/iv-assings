class Solution {
    public int lengthOfLongestSubstring(String str) {
        Set<Character> s = new HashSet<Character>();
        int maxLen = 0;
        int leftmost = 0;
        for(int i = 0; i < str.length(); i++) {
            while(s.contains(str.charAt(i))) {
                s.remove(str.charAt(leftmost++));
            }
            s.add(str.charAt(i));
            maxLen=Math.max(maxLen, (int)s.size());
        }
        return maxLen; 
    }
}
