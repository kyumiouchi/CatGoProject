using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ParticleClick : MonoBehaviour
{
	[SerializeField] ParticleSystem ps;

	Queue<ParticleSystem> _queue = new Queue<ParticleSystem>();

    void Update()
    {
        if (Input.GetMouseButton(0))
        {
            Vector3 clickPosition = Camera.main.ScreenToWorldPoint(Input.mousePosition);
            SpawnParticleAt(clickPosition);
        }
    }

    private void SpawnParticleAt(Vector3 clickPosition)
    {
        if (_queue.Count == 0 || _queue.Peek().isPlaying)
        {
            var instantiate = Instantiate(ps, clickPosition, Quaternion.identity);
            instantiate.Play();
            instantiate.name = "Particle_" + _queue.Count;
            instantiate.transform.parent = transform;
            _queue.Enqueue(instantiate);
        }
        else
        {
            var instance = _queue.Dequeue();
            instance.gameObject.SetActive(true);
            instance.transform.position = clickPosition;
            instance.Play();
            _queue.Enqueue(instance);
        }
    }
}
