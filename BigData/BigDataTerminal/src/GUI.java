
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;
import java.net.URI;
import java.net.UnknownHostException;

import javax.swing.*;

class Client {

    private final PrintWriter cout;
    private final BufferedReader br;

    public Client(String addr) throws IOException,
            InterruptedException {
        String[] hostPort = addr.split(":");
        Socket socket = new Socket(hostPort[0], Integer.parseInt(hostPort[1]));   //连接服务器9000端口

        this.cout = new PrintWriter(socket.getOutputStream());

        this.br = new BufferedReader(new InputStreamReader(
                socket.getInputStream()));
    }

    public void send(String text) throws IOException {
        this.cout.print(text);     //向服务器写入数据
        this.cout.flush();

    }

    public String read() throws IOException {
        char[] rev = new char[100];
        int n = this.br.read(rev);     //读取服务器数据
        StringBuilder res = new StringBuilder();
        for (int i = 0; i < n; i++) {
            res.append(rev[i]);
        }
        System.out.println(res);
        return res.toString();
    }
}


public class GUI extends JFrame{
    Client client;
    public GUI(){//定义一个构造方法
        inputPanel();
    }

    public void inputPanel(){
        JFrame frame=new JFrame("Input");    //创建Frame窗口
        frame.setSize(600,200);
        JPanel jp=new JPanel();    //创建面板
        JTextField txtfield=new JTextField(30);
        txtfield.setHorizontalAlignment(JTextField.CENTER);
        JButton button = new JButton("ok");
        button.addActionListener(new ActionListener() {
            public void actionPerformed(ActionEvent e) {
                String s = txtfield.getText();
                try {
                    client = new Client(s);
                    frame.setVisible(true);
                    selectPanel();
                } catch (IOException | InterruptedException ioException) {
                    ioException.printStackTrace();
                }
            }
        });
        jp.add(txtfield);
        jp.add(button);
        frame.add(jp);
        frame.setVisible(true);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
    }

    public void selectPanel() throws IOException, InterruptedException {
        //获取图片所在的URL    以下2行代码需要额外注意
        Icon jupyterImage=new ImageIcon(Toolkit.getDefaultToolkit().getImage(this.getClass().getResource("jupyter.png")));//实例化Icon对象
        Icon sparkImage=new ImageIcon(Toolkit.getDefaultToolkit().getImage(this.getClass().getResource("spark.png")));//实例化Icon对象
        Icon hadoopImage=new ImageIcon(Toolkit.getDefaultToolkit().getImage(this.getClass().getResource("hadoop.png")));//实例化Icon对象
        Icon sonarqubeImage=new ImageIcon(Toolkit.getDefaultToolkit().getImage(this.getClass().getResource("sonarqube.png")));//实例化Icon对象


        //设置网格布局管理器   3行2列  水平5垂直5
        setLayout(new GridLayout(2,2,5,5));
        Container container=getContentPane();

        JButton jupyterButton=new JButton("",jupyterImage);
        JButton sparkButton=new JButton("",sparkImage);
        JButton hadoopButton=new JButton("",hadoopImage);
        JButton sonarqubeButton=new JButton("",sonarqubeImage);


        hadoopButton.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                try {
                    client.send("1");
                    browse(client.read());
                } catch (Exception ioException) {
                    ioException.printStackTrace();
                }
            }
        });

        sparkButton.addActionListener(new ActionListener() {

            @Override
            public void actionPerformed(ActionEvent e) {
                try {
                    client.send("2");
                    browse(client.read());
                } catch (Exception ioException) {
                    ioException.printStackTrace();
                }
            }
        });

        jupyterButton.addActionListener(new ActionListener() {

            @Override
            public void actionPerformed(ActionEvent e) {
                try {
                    client.send("3");
                    browse(client.read());
                } catch (Exception ioException) {
                    ioException.printStackTrace();
                }
            }
        });

        sonarqubeButton.addActionListener(new ActionListener() {

            @Override
            public void actionPerformed(ActionEvent e) {
                try {
                    client.send("4");
                    browse(client.read());
                } catch (Exception ioException) {
                    ioException.printStackTrace();
                }
            }
        });


        container.add(jupyterButton);//将按钮添加到容器中
        container.add(sparkButton);//将按钮添加到容器中
        container.add(hadoopButton);//将按钮添加到容器中
        container.add(sonarqubeButton);//将按钮添加到容器中

        setTitle("BigDataTerminal");//设置窗口标题
        setVisible(true);//设置窗口可视化
        setSize(1000,1000);//设置窗口的大小
        //设置窗口的关闭方式
        setDefaultCloseOperation(WindowConstants.EXIT_ON_CLOSE);
    }

    private static void browse(String url) throws Exception {
        Desktop desktop = Desktop.getDesktop();
        if (Desktop.isDesktopSupported() && desktop.isSupported(Desktop.Action.BROWSE)) {
            URI uri = new URI(url);
            desktop.browse(uri);
        }
    }

    public static void main(String[] args) throws Exception {
        // TODO Auto-generated method stub
        GUI jb=new GUI();
    }

}