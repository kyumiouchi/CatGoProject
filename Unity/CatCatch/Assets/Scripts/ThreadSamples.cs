using System.Collections;
using System.Collections.Generic;
using System.Threading;
using UnityEngine;

public class ThreadSamples : MonoBehaviour
{
	public float[] randomList;
    // Start is called before the first frame update
    void Start()
    {
	    randomList = new float[1_000_000_000];
	    Thread t = new Thread(delegate()
	    {
		    while (true)
		    {
			    Generate();
			    Thread.Sleep(16);
		    }
	    });
	    t.Start();
    }

    private void Generate()
    {
	    System.Random value = new System.Random();
	    for (int i = 0; i < randomList.Length; i++)
	    {
		    randomList[i] = (float)value.NextDouble();
	    }
	    Debug.Log("END");
    }
}
