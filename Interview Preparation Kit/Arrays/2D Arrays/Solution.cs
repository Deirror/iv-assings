using System.CodeDom.Compiler;
using System.Collections.Generic;
using System.Collections;
using System.ComponentModel;
using System.Diagnostics.CodeAnalysis;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Runtime.Serialization;
using System.Text.RegularExpressions;
using System.Text;
using System;

class Result
{

    /*
     * Complete the 'hourglassSum' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts 2D_INTEGER_ARRAY arr as parameter.
     */

    public static int hourglassSum(List<List<int>> arr)
    {
        int lrow = arr.Count - 1,
            lcol = arr[0].Count - 1;
        
        int result = Int32.MinValue;
        
        for(int i = 1; i < lcol; i++)
        {
            for (int j = 1; j < lrow; j++)
            {
                int firstRow = arr[i - 1][j - 1] + arr[i - 1][j] + arr[i - 1][j + 1];
                int thirdRow = arr[i + 1][j - 1] + arr[i + 1][j] + arr[i + 1][j + 1];
                
                int sum = firstRow + arr[i][j] + thirdRow;
                
                result = Int32.Max(result, sum);
            }
        }
        
        return result;
    }

}

class Solution
{
    public static void Main(string[] args)
    {
        TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        List<List<int>> arr = new List<List<int>>();

        for (int i = 0; i < 6; i++)
        {
            arr.Add(Console.ReadLine().TrimEnd().Split(' ').ToList().Select(arrTemp => Convert.ToInt32(arrTemp)).ToList());
        }

        int result = Result.hourglassSum(arr);

        textWriter.WriteLine(result);

        textWriter.Flush();
        textWriter.Close();
    }
}
