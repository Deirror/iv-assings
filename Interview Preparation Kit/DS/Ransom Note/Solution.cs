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
     * Complete the 'checkMagazine' function below.
     *
     * The function accepts following parameters:
     *  1. STRING_ARRAY magazine
     *  2. STRING_ARRAY note
     */

    public static void checkMagazine(List<string> magazine, List<string> note)
    {
        IDictionary<string, int> magazineHistogram = new Dictionary<string, int>();
        IDictionary<string, int> noteHistogram = new Dictionary<string, int>();
        
        foreach(string s in magazine)
        {
            if(magazineHistogram.ContainsKey(s))
            {
                magazineHistogram[s]++;
            }
            else
            {
                magazineHistogram.Add(s, 1);
            }
        }
        
        foreach(string s in note)
        {
            if(noteHistogram.ContainsKey(s))
            {
                noteHistogram[s]++;
            }
            else
            {
                noteHistogram.Add(s, 1);
            }
        }
        
        foreach(var kvp in noteHistogram)
        {
            if(magazineHistogram.ContainsKey(kvp.Key))
            {
                if(magazineHistogram[kvp.Key] < kvp.Value)
                {
                    Console.WriteLine("No");
                    return;
                }
            }
            else
            {
                Console.WriteLine("No");
                return;
            }
        }
        
        Console.WriteLine("Yes");
    }

}

class Solution
{
    public static void Main(string[] args)
    {
        string[] firstMultipleInput = Console.ReadLine().TrimEnd().Split(' ');

        int m = Convert.ToInt32(firstMultipleInput[0]);

        int n = Convert.ToInt32(firstMultipleInput[1]);

        List<string> magazine = Console.ReadLine().TrimEnd().Split(' ').ToList();

        List<string> note = Console.ReadLine().TrimEnd().Split(' ').ToList();

        Result.checkMagazine(magazine, note);
    }
}
