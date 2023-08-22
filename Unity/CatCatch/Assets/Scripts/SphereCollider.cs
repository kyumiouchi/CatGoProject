using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class SphereCollider : MonoBehaviour

{
    private void Start()
    {
        transform.position = new Vector3(10f, transform.position.y, transform.position.z);
    }

    private void OnTriggerEnter(Collider other)
    {
        Debug.Log($"Collider Name {other.name}");
    }
}
