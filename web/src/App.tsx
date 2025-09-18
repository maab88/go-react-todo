import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useState } from 'react'
import { api } from './api/client'

type Task = {
  id: number
  title: string
  description: string
  dueDate?: string | null
  priority: number
  status: string
  createdAt: string
  updatedAt: string
}

export default function App() {
  const qc = useQueryClient()
  const [title, setTitle] = useState('')
  const [q, setQ] = useState('')
  const [status, setStatus] = useState('')

  const { data, isLoading } = useQuery({
    queryKey: ['tasks', { q, status }],
    queryFn: async () => {
      const params = new URLSearchParams()
      if (q) params.set('q', q)
      if (status) params.set('status', status)
      const res = await api.get(`/tasks/?${params.toString()}`)
      return res.data.data as Task[]
    }
  })

  const create = useMutation({
    mutationFn: async (title: string) => {
      const res = await api.post('/tasks', { title })
      return res.data.data as Task
    },
    onSuccess: () => {
      setTitle('')
      qc.invalidateQueries({ queryKey: ['tasks'] })
    }
  })

  const toggle = useMutation({
    mutationFn: async (t: Task) => {
      const newStatus = t.status === 'done' ? 'todo' : 'done'
      await api.patch(`/tasks/${t.id}/status`, { status: newStatus })
    },
    onSuccess: () => qc.invalidateQueries({ queryKey: ['tasks'] })
  })

  const remove = useMutation({
    mutationFn: async (id: number) => {
      await api.delete(`/tasks/${id}`)
    },
    onSuccess: () => qc.invalidateQueries({ queryKey: ['tasks'] })
  })

  return (
    <div className="max-w-3xl mx-auto p-6">
      <h1 className="text-3xl font-bold mb-4">Go + React Todo</h1>

      <div className="flex gap-2 mb-4">
        <input
          value={title}
          onChange={e => setTitle(e.target.value)}
          placeholder="Add a task and press Enter"
          onKeyDown={e => {
            if (e.key === 'Enter' && title.trim()) create.mutate(title.trim())
          }}
          className="flex-1 px-3 py-2 border rounded-lg"
        />
        <button
          onClick={() => title.trim() && create.mutate(title.trim())}
          className="px-4 py-2 rounded-lg bg-black text-white"
        >
          Add
        </button>
      </div>

      <div className="flex gap-2 mb-4">
        <input
          value={q}
          onChange={e => setQ(e.target.value)}
          placeholder="Search"
          className="flex-1 px-3 py-2 border rounded-lg"
        />
        <select value={status} onChange={e => setStatus(e.target.value)} className="px-3 py-2 border rounded-lg">
          <option value="">All</option>
          <option value="todo">Todo</option>
          <option value="done">Done</option>
        </select>
      </div>

      {isLoading ? (
        <p>Loading…</p>
      ) : (
        <ul className="space-y-2">
          {data?.map(t => (
            <li key={t.id} className="flex items-center justify-between p-3 rounded-xl bg-white shadow">
              <div className="flex items-center gap-3">
                <input
                  type="checkbox"
                  checked={t.status === 'done'}
                  onChange={() => toggle.mutate(t)}
                />
                <span className={t.status === 'done' ? 'line-through text-gray-500' : ''}>{t.title}</span>
              </div>
              <button onClick={() => remove.mutate(t.id)} className="text-red-600">Delete</button>
            </li>
          ))}
          {!data?.length && <p className="text-gray-500">No tasks yet.</p>}
        </ul>
      )}
    </div>
  )
}
