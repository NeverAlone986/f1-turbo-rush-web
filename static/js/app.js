async function changeTrack(trackNum) {
    try {
        const response = await fetch(`/api/track/${trackNum}`, {
            method: 'POST'
        });
        if (response.ok) {
            alert(`Track ${trackNum} selected!`);
        }
    } catch (error) {
        console.error('Error changing track:', error);
    }
}

// Пример обновления состояния игры
async function updateGameState() {
    try {
        const response = await fetch('/api/state');
        const state = await response.json();
        document.getElementById('score').textContent = state.Score;
        document.getElementById('record').textContent = state.Record;
    } catch (error) {
        console.error('Error fetching game state:', error);
    }
}

// Обновляем состояние каждую секунду
setInterval(updateGameState, 1000);
