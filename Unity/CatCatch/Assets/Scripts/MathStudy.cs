using System;
using UnityEngine;

public class MathStudy : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
        ExecuteExamples();
        //ExecuteMath();
    }

    private void ExecuteMath()
    {
        Debug.Log("Pow " + Math.Pow(6,3));
    }

    private void ExecuteExamples()
    {
        InRoundNumber(6.283f);
    }

    void InRoundNumber(float number)
    {
        // Round
        Debug.Log($"Round value: {Mathf.Round(number)}");
        Debug.Log($"Round value 1dp: {Math.Round(number, 1)}");
        Debug.Log($"Round value 2dp: {Math.Round(number, 2)}");
        Debug.Log($"Ceil value: {Mathf.Ceil(number)}");
        Debug.Log($"Floor value: {Mathf.Floor(number)}");
    }
}
