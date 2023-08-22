using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class MoveCubeToPoint : MonoBehaviour
{
    public float speed;
    public GameObject target; 
    void Update()
    {
        float distance = Vector3.Distance(this.transform.position, target.transform.position);
        if (distance > 1f)
        {
            Vector3 disVect = target.transform.position - transform.position;
            disVect.Normalize();
            transform.position += Time.deltaTime * speed * disVect;
        }
    }
}
