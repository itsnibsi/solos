import { useState, useEffect } from 'react'

const API_URL = import.meta.env.PROD ? '/api' : '/api'

function App() {
  const [health, setHealth] = useState<string>('Loading...')

  useEffect(() => {
    fetch(`${API_URL}/health`)
      .then((response) => response.json())
      .then((data) => setHealth(data.status))
      .catch(() => setHealth('Error'))
  }, [])

  return (
    <div>
      <h1>Solos</h1>
      <p>API Health: {health}</p>
    </div>
  )
}

export default App
