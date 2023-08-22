using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class SpawnEnvironment : MonoBehaviour
{
    [SerializeField] private Transform environment;
    
    void Update()
    {
        if (Input.GetKeyDown(KeyCode.T))
        {
            Instantiate(environment);
        }
    }
}
