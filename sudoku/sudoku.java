import java.util.*;
import java.io.*;
import java.math.*;
import java.util.stream.Collectors;
/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
class Solution {
    final int[][] cells;

    Solution(int[][] cells) {
        this.cells = cells;
    }

    boolean isValid() {
        List<Set<Integer>> sets = new ArrayList<>();
        for(int i =0; i < 27; i++) {
            sets.add(new TreeSet<>());
        }
        for(int i = 0; i < 9; i++) {
            for(int j=0; j < 9; j++) {
                int sqId = (((j%9) / 3 + (i/3) * 3));
                sets.get(i).add(cells[i][j]);
                sets.get(i+9).add(cells[j][i]);
                sets.get(sqId+18).add(cells[i][j]);
            }
        }
        return sets.stream().filter(s -> s.size() == 9 ).collect(Collectors.toList()).size() == 27;
    }

    public static void main(String args[]) {
        Scanner in = new Scanner(System.in);
        int[][] cells = new int[9][9];
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                cells[i][j]=in.nextInt();
            }
        }
        Solution s = new Solution(cells);
        String result = s.isValid() ? "true" : "false";
        // Write an answer using System.out.println()
        // To debug: System.err.println("Debug messages...");
        System.out.println(result);
    }
}
