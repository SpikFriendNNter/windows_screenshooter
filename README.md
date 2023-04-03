# <strong>Como é feito o processo de screenshot utilizando a API do Windows?</strong>

<img src="https://assets.labs.ine.com/web/badges/low/winapi.png" width=200>

## <strong>Processo</strong>

### O processo de captura de imagens utilizando a api do windows se dá pela seguinte estrutura:

- Obter o DC(Device Context) da tela.
    
- Criar um objeto Bitmap compatível com o DC de origem.
    
- Duplicar esse DC, criando um novo DC compatível com o original.
    
- É preciso definir alguns valores padrões que um header bitmap carrega.
    
- Após a criação do objeto bitmap, é preciso alocar espaço suficiente para conter o tamanho do bitmap na memória do sistema operacional para que sirva de buffer, e então, podermos acessar os dados do bitmap, que posteriormente, conterá os dados de cores obtidos pelo DC de origem.

- Utilizamos funções da API do windows para que seja feita a transferência/copia dos dados de cores do DC de origem para o DC de destino. O objeto Bitmap dentro do DC de destino irá conter os dados de cores obtidos do DC de origem. Normalmente, a função que realiza tal tarefa é BitBlt da dll gdi32.

- Após isso, usamos a função GetDIBits que irá transferir os dados do objeto Bitmap no DC de destino para o buffer que alocamos na memoria.

- Com isso feito, podemos obter os dados de cores RGB a partir do buffer como um rolo de bits(como um array), que tem a seguinte ordem de dados: BGR => Blue, Green, Red.

- Para finalizar, lembre-se de liberar o buffer alocado em memória, liberar os DCs e window Handlers, e fechar os objetos bitmap.

## <strong>O que é DC?</strong>
- Segundo a documentação do windows (https://learn.microsoft.com/en-us/cpp/mfc/device-contexts?view=msvc-170), um DC ou Device Context é uma estrutura de dados do Windows contendo informações sobre atributos de desenho de um dispositivo como um monitor ou impressora. Em resumo, é uma estrutura de dados para representação de um dispositivo gráfico. Para mais detalhes, consulte a documentação indicada no link acima.

## <strong>Como obter as especificações de tela?</strong>

- <span style="color:yellow;">GetMonitorInfoW:</span> função que obtém as especificações do monitor especificado. Lembrando-que, para que esta função funcione, é necessário fornecer um parâmetro, que é um <i style="color:#9dcec7;">HANDLE</i> para o monitor da janela especificada. <strong style="color:red;">Atenção:</strong> Passar como parâmetro o HANDLE da janela especificada causará um erro. O correto é obter o <i style="color:#9dcec7;">HANDLE</i> do monitor utilizando a função MonitorFromWindow.

- <span style="color:yellow;">MonitorFromWindow</span>: função que recebe como parâmetro um <i style="color:#9dcec7;">HANDLE</i> para uma janela especificada e retorna um <i style="color:#9dcec7;">HANDLE</i> de monitor daquela janela.

## <strong>Funções da API do Windows utilizadas neste projeto</strong>
- <span style="color:yellow;">GetDesktopWindow</span>;
- <span style="color:yellow;">MonitorFromWindow</span>;
- <span style="color:yellow;">GetMonitorInfoW</span>;
- <span style="color:yellow;">GetDC</span>;
- <span style="color:yellow;">ReleaseDC</span>;
- <span style="color:yellow;">CreateCompatibleDC</span>;
- <span style="color:yellow;">CreateCompatibleBitmap</span>;
- <span style="color:yellow;">SelectObject</span>;
- <span style="color:yellow;">BitBlt</span>;
- <span style="color:yellow;">GetDIBits</span>;
- <span style="color:yellow;">GlobalAlloc</span>;
- <span style="color:yellow;">GlobalFree</span>;
- <span style="color:yellow;">GlobalLock</span>;

## Projeto utilizado como base
- https://github.com/kbinani/screenshot