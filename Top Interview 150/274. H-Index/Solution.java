class Solution {
    public int hIndex(int[] citations) {
        Integer[] what = Arrays.stream(citations).boxed().toArray( Integer[]::new );
        Arrays.sort(what, Collections.reverseOrder());
        int h = 0;
        for(int i = 0; i < what.length; i++) {
            if(what[i] >= i + 1) {
                h = i + 1;
            } else {
                break;
            }
        }
        return h;
    }
}
