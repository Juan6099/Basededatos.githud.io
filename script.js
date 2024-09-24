document.getElementById('mostrarEjemplo').addEventListener('click', () => {
    Swal.fire({
        title: 'Ejemplo de Base de Datos: Empleados',
        html: `
            <table class="table-auto border-collapse border border-gray-400">
                <thead>
                    <tr>
                        <th class="border border-gray-400 p-2">ID</th>
                        <th class="border border-gray-400 p-2">Nombre</th>
                        <th class="border border-gray-400 p-2">Departamento</th>
                        <th class="border border-gray-400 p-2">Salario</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td class="border border-gray-400 p-2">1</td>
                        <td class="border border-gray-400 p-2">Juan Pérez</td>
                        <td class="border border-gray-400 p-2">Ventas</td>
                        <td class="border border-gray-400 p-2">3000</td>
                    </tr>
                    <tr>
                        <td class="border border-gray-400 p-2">2</td>
                        <td class="border border-gray-400 p-2">María Gómez</td>
                        <td class="border border-gray-400 p-2">Contabilidad</td>
                        <td class="border border-gray-400 p-2">2500</td>
                    </tr>
                    </tbody>
            </table>
        `,
        showConfirmButton: false,
    });
});