document.addEventListener('DOMContentLoaded', () => {
    const personForm = document.getElementById('personForm');
    const nameInput = document.getElementById('name');
    const ageInput = document.getElementById('age');
    const btnSubmit = document.getElementById('btnSubmit');
    const personTableBody = document.querySelector('#personTable tbody');

    const statusModal = document.getElementById('statusModal');
    const modalCard = document.querySelector('.modal-card');
    const modalTitle = document.getElementById('modalTitle');
    const modalMessage = document.getElementById('modalMessage');
    const btnCloseModal = document.getElementById('btnCloseModal');

    let persons = [];

    personForm.addEventListener('submit', async (e) => {
        e.preventDefault();

        const rawName = nameInput.value.trim();
        const rawAge = parseInt(ageInput.value, 10);

        if (!rawName || isNaN(rawAge)) return;

        btnSubmit.disabled = true;
        btnSubmit.textContent = 'Registering...';

        try {
            const response = await fetch('/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    RawName: rawName,
                    RawAge: rawAge
                })
            });

            const result = await response.json();

            if (response.ok && result.success) {
                showModal({
                    type: 'success',
                    title: 'Berhasil',
                    message: result.message || 'Person berhasil didaftarkan!'
                });
                
                persons.unshift(result.data);
                renderTable();

                nameInput.value = '';
                ageInput.value = '';
            } else {
                showModal({
                    type: 'error',
                    title: 'Gagal',
                    message: result.message || 'Terjadi kesalahan pada sistem.'
                });
            }
        } catch (error) {
            showModal({
                type: 'error',
                title: 'Koneksi Gagal',
                message: 'Gagal terhubung ke server backend.'
            });
        } finally {
            btnSubmit.disabled = false;
            btnSubmit.textContent = 'Register Person';
        }
    });

    window.deletePerson = async function(id) {
        try {
            const response = await fetch(`/person?id=${id}`, {
                method: 'DELETE'
            });

            const result = await response.json();

            if (response.ok && result.success) {
                showModal({
                    type: 'success',
                    title: 'Berhasil',
                    message: result.message || 'Person berhasil dihapus!'
                });

                persons = persons.filter(p => p.id !== id);
                renderTable();
            } else {
                showModal({
                    type: 'error',
                    title: 'Gagal',
                    message: result.message || 'Gagal menghapus person.'
                });
            }
        } catch (error) {
            showModal({
                type: 'error',
                title: 'Koneksi Gagal',
                message: 'Gagal terhubung ke server backend.'
            });
        }
    };

    function renderTable() {
        if (persons.length === 0) {
            personTableBody.innerHTML = `
                <tr class="empty-row">
                    <td colspan="4">No data available</td>
                </tr>
            `;
            return;
        }

        personTableBody.innerHTML = persons.map(p => `
            <tr>
                <td><code>${p.id}</code></td>
                <td>${p.name}</td>
                <td>${p.age}</td>
                <td>
                    <button class="btn-delete" onclick="deletePerson('${p.id}')">Hapus</button>
                </td>
            </tr>
        `).join('');
    }

    function showModal({ type, title, message }) {
        modalCard.className = 'modal-card ' + type;
        modalTitle.textContent = title;
        modalMessage.textContent = message;
        statusModal.classList.remove('hidden');
    }

    function closeModal() {
        statusModal.classList.add('hidden');
    }

    btnCloseModal.addEventListener('click', closeModal);

    statusModal.addEventListener('click', (e) => {
        if (e.target === statusModal) {
            closeModal();
        }
    });
});