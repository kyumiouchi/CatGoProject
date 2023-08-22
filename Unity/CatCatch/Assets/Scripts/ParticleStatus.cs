using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ParticleStatus : MonoBehaviour
{
    [SerializeField] private string status;
    private ParticleSystem _particleSystem;


    private void Start()
    {
        _particleSystem = GetComponent<ParticleSystem>();
    }


    void Update()
    {
        if (_particleSystem.isPlaying)
        {
            status = "isPlaying";
        } else if (_particleSystem.isEmitting)
        {
            status = "isEmitting";
        } else if (_particleSystem.isPaused)
        {
            status = "isPaused";
        } else if (_particleSystem.isStopped)
        {
            status = "isStopped";
        } 
    }
}
