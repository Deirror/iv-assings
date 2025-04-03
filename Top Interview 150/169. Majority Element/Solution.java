class Solution {
    public int majorityElement(int[] nums) {
        Map<Integer, Integer> histogram = new HashMap<Integer, Integer>();
        for(int i = 0; i < nums.length; i++) {
            histogram.put(nums[i], histogram.getOrDefault(nums[i], 0) + 1);
        }
        int time = nums.length / 2;
        int majorityElement = 0;
        for(var entry : histogram.entrySet()) {
            if(entry.getValue() > time) {
                majorityElement = entry.getKey();
            }
        }
        return majorityElement;
    }
}
