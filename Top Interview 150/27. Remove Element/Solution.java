class Solution {
    public int removeElement(int[] nums, int val) {
        int k = 0;
        for(int i = 0; i < nums.length; i++) {
            if(nums[i] != val) {
                if(i == k) {
                    k++;
                    continue;
                }
                int buff = nums[i];
                nums[i] = nums[k];
                nums[k] = buff;
                k++;
            }
        }
        return k;
    }
}
