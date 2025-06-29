import { writable } from 'svelte/store';
import Cookies from 'js-cookie';

export function createCmdStore(id: string) {
    const output = writable<string[]>([]);
    let socket: WebSocket;
    let token = Cookies.get("token");
    let connected = false;

    function sendCommand(command: string) {
        if (socket && connected && command.trim() !== "") {
            socket.send(command.trim());
        }
    }

    function startWebSocket() {
        if (!id || !token) {
            output.set(["❌ Container ID veya kullanıcı token eksik."]);
            return;
        }

        const wsUrl = `ws://localhost:3000/ws/${id}?token=${token}`;
        socket = new WebSocket(wsUrl);

        socket.onopen = () => {
            connected = true;
            output.update(lines => [...lines, "✅ Bağlantı kuruldu."]);
        };

        socket.onmessage = (event) => {
            const msg = event.data;
            if (msg === "__CLEAR__") {
                output.set([]);
            } else {
                output.update(lines => [...lines, msg]);
            }
        };

        socket.onerror = (err) => {
            output.update(lines => [...lines, "❌ WebSocket hatası: " + err]);
        };

        socket.onclose = () => {
            connected = false;
            output.update(lines => [...lines, "🔌 Bağlantı kapandı."]);
        };
    }

    return {
        output,
        sendCommand,
        startWebSocket
    };
}
