document.addEventListener('DOMContentLoaded', () => {
    const emailForm = document.getElementById('emailForm');
    const receiverInput = document.getElementById('receiver');
    const btnSend = document.getElementById('btnSend');

    // Element Modal
    const statusModal = document.getElementById('statusModal');
    const modalCard = document.querySelector('.modal-card');
    const modalTitle = document.getElementById('modalTitle');
    const modalMessage = document.getElementById('modalMessage');
    const btnCloseModal = document.getElementById('btnCloseModal');

    // Event Handling Form Submission
    emailForm.addEventListener('submit', async (e) => {
        e.preventDefault();

        const receiverEmail = receiverInput.value.trim();
        if (!receiverEmail) return;

        // Set state loading pada tombol
        btnSend.disabled = true;
        btnSend.textContent = 'Sending...';

        try {
            // Mengirim request POST ke endpoint backend .NET
            const response = await fetch('/api/email/send', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    receiver: receiverEmail
                })
            });

            // Parsing JSON dari respons controller .NET
            const result = await response.json();

            if (response.ok) {
                // Tampilkan Pop-up Sukses dengan tema semantik hijau
                showModal({
                    type: 'success',
                    title: 'Berhasil',
                    message: result.message || 'Email berhasil dikirim!'
                });
                receiverInput.value = ''; // Reset form input
            } else {
                // Tampilkan Pop-up Gagal dengan tema semantik merah
                showModal({
                    type: 'error',
                    title: 'Gagal Terkirim',
                    message: result.message || 'Terjadi kesalahan pada sistem.'
                });
            }
        } catch (error) {
            // Menangani error jaringan (network / CORS)
            showModal({
                type: 'error',
                title: 'Koneksi Gagal',
                message: error.message || 'Gagal terhubung ke server backend.'
            });
        } finally {
            // Kembalikan state tombol ke semula
            btnSend.disabled = false;
            btnSend.textContent = 'Send Email';
        }
    });

    // Fungsi Menampilkan Modal Pop-up
    function showModal({ type, title, message }) {
        // Terapkan kelas semantik ('success' atau 'error') ke elemen modal-card
        modalCard.className = 'modal-card ' + type;

        modalTitle.textContent = title;
        modalMessage.textContent = message;
        statusModal.classList.remove('hidden');
    }

    // Fungsi Menutup Modal
    function closeModal() {
        statusModal.classList.add('hidden');
    }

    // Event Menutup Modal lewat Tombol Tutup
    btnCloseModal.addEventListener('click', closeModal);

    // Tutup modal jika pengguna mengklik area luar (overlay) card
    statusModal.addEventListener('click', (e) => {
        if (e.target === statusModal) {
            closeModal();
        }
    });
});